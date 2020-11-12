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

func run(pass *analysis.Pass) (interface{}, error) {
	ssaInfo := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	ssaFuncMap := getSSAFuncMap(ssaInfo)

	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			if funcDecl, _ := decl.(*ast.FuncDecl); funcDecl != nil {
				loc := measure.LineOfCode(pass.Fset, funcDecl)
				maxnesting := measure.MaxNestingLevel(funcDecl)
				noav := measure.NumberOfAccessedVariables(funcDecl, pass.TypesInfo)
				ssaFunc, _ := ssaFuncMap[funcDecl.Name.Name]
				cyclo := measure.CyclomaticComplexity(ssaFunc)
				fmt.Println(funcDecl.Name.Name, loc, maxnesting, noav, cyclo)
			}
		}
	}

	return nil, nil
}

func getSSAFuncMap(ssaInfo *buildssa.SSA) map[string]*ssa.Function {
	ssaFuncMap := map[string]*ssa.Function{}

	for _, ssaFunc := range ssaInfo.SrcFuncs {
		ssaFuncMap[ssaFunc.Name()] = ssaFunc
	}

	return ssaFuncMap
}
