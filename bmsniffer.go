package bmsniffer

import (
	"fmt"
	"go/ast"

	"github.com/oribe1115/bmsniffer/measure"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ssa"
)

const doc = "bmsniffer is ..."

var (
	locLimit        int
	maxnestingLimit int
	novLimit        int
	cycloLimit      int
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

type FuncData struct {
	FuncDecl   *ast.FuncDecl
	SsaFunc    *ssa.Function
	Loc        int
	Maxnesting int
	Nov        int
	Cyclo      int
}

type FuncDataList []*FuncData

func init() {
	Analyzer.Flags.IntVar(&locLimit, "loc", 0, "limit for LOC")
	Analyzer.Flags.IntVar(&maxnestingLimit, "maxnesting", 0, "limit for MAXNESTING")
	Analyzer.Flags.IntVar(&novLimit, "nov", 0, "limit for NOV")
	Analyzer.Flags.IntVar(&cycloLimit, "cyclo", 0, "limit for CYCLO")
}

func run(pass *analysis.Pass) (interface{}, error) {
	ssaInfo := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	ssaFuncMap := getSSAFuncMap(ssaInfo)

	list := &FuncDataList{}

	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			if funcDecl, _ := decl.(*ast.FuncDecl); funcDecl != nil {
				ssaFunc, _ := ssaFuncMap[funcDecl.Name.Name]
				funcData := &FuncData{
					FuncDecl:   funcDecl,
					SsaFunc:    ssaFunc,
					Loc:        measure.LineOfCode(pass.Fset, funcDecl),
					Maxnesting: measure.MaxNestingLevel(funcDecl),
					Nov:        measure.NumberOfAccessedVariables(funcDecl, pass.TypesInfo),
					Cyclo:      measure.CyclomaticComplexity(ssaFunc),
				}
				list.Add(funcData)
			}
		}
	}

	filterdList := list.Filterd(func(fd *FuncData) bool {
		return fd.Loc >= locLimit && fd.Maxnesting >= maxnestingLimit && fd.Nov >= novLimit && fd.Cyclo >= cycloLimit
	})

	filterdList.PrintAll()

	return nil, nil
}

func getSSAFuncMap(ssaInfo *buildssa.SSA) map[string]*ssa.Function {
	ssaFuncMap := map[string]*ssa.Function{}

	for _, ssaFunc := range ssaInfo.SrcFuncs {
		ssaFuncMap[ssaFunc.Name()] = ssaFunc
	}

	return ssaFuncMap
}

func (fl *FuncDataList) Add(funcData *FuncData) {
	*fl = append(*fl, funcData)
}

func (fl *FuncDataList) Filterd(filter func(*FuncData) bool) *FuncDataList {
	newList := &FuncDataList{}
	for _, funcData := range *fl {
		if filter(funcData) {
			newList.Add(funcData)
		}
	}

	return newList
}

func (fl *FuncDataList) PrintAll() {
	fmt.Println("funcName: LOC-MAXNESTING-NOV-CYCLO")
	for _, funcData := range *fl {
		fmt.Println(funcData.FuncDecl.Name.String(), funcData.Loc, funcData.Maxnesting, funcData.Nov, funcData.Cyclo)
	}
}
