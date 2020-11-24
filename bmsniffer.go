package bmsniffer

import (
	"go/ast"
	"regexp"

	"github.com/oribe1115/bmsniffer/analyzed"
	"github.com/oribe1115/bmsniffer/measure"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "bmsniffer is ..."

var (
	locLimit        int
	maxnestingLimit int
	noavLimit       int
	cycloLimit      int
	includeTest     bool
)

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "bmsniffer",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
		buildssa.Analyzer,
	},
}

func init() {
	Analyzer.Flags.IntVar(&locLimit, "loc", 0, "baseline for LOC")
	Analyzer.Flags.IntVar(&maxnestingLimit, "maxnesting", 0, "baseline for MAXNESTING")
	Analyzer.Flags.IntVar(&noavLimit, "noav", 0, "baseline for NOAV")
	Analyzer.Flags.IntVar(&cycloLimit, "cyclo", 0, "baseline for CYCLO")
	Analyzer.Flags.BoolVar(&includeTest, "test", false, "include test codes for analysis")
}

func run(pass *analysis.Pass) (interface{}, error) {
	ssaInfo := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	ssaData := measure.GetSSAData(ssaInfo)
	fset := ssaInfo.Pkg.Prog.Fset

	pkgData := analyzed.NewPkg()

	for _, file := range pass.Files {
		fileName := fset.File(file.Pos()).Name()
		fileData := analyzed.NewFile(fileName, file)

		for _, decl := range file.Decls {
			if funcDecl, _ := decl.(*ast.FuncDecl); funcDecl != nil {
				funcData := &analyzed.Func{
					FuncDecl:   funcDecl,
					Loc:        measure.LineOfCode(pass.Fset, funcDecl),
					Maxnesting: measure.MaxNestingLevel(funcDecl),
					Noav:       measure.NumberOfAccessedVariables(funcDecl, pass.TypesInfo),
					Cyclo:      measure.CyclomaticComplexity(funcDecl.Name.Name, ssaData),
				}
				fileData.AddFunc(funcData)
			}
		}

		pkgData.AddFile(fileData)
	}

	pkgData.AddFileFilter(func(file *analyzed.File) bool {
		testFileRegExp := regexp.MustCompile(`.*_test\.go$`)
		if !includeTest && testFileRegExp.MatchString(file.Name) {
			return false
		}
		return true
	})

	pkgData.AddFuncFilter(func(fn *analyzed.Func) bool {
		return fn.Loc >= locLimit && fn.Maxnesting >= maxnestingLimit && fn.Noav >= noavLimit && fn.Cyclo >= cycloLimit
	})

	pkgData.Filter()
	pkgData.Print()

	return nil, nil
}
