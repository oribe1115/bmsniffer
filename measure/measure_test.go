package measure

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ssa"
)

func getFsetAndFuncDecl(t *testing.T, filename string) (*token.FileSet, *ast.FuncDecl) {
	t.Helper()

	testdata := analysistest.TestData()
	result := analysistest.Run(t, testdata, inspect.Analyzer, filename)[0]
	pass := result.Pass

	var funcDecl *ast.FuncDecl
	for _, decl := range pass.Files[0].Decls {
		funcDecl, _ = decl.(*ast.FuncDecl)
		if funcDecl != nil {
			break
		}
	}

	return pass.Fset, funcDecl
}

func getFuncDeclAndTypeInfo(t *testing.T, filename string) (*ast.FuncDecl, *types.Info) {
	t.Helper()

	testdata := analysistest.TestData()
	result := analysistest.Run(t, testdata, inspect.Analyzer, filename)[0]
	pass := result.Pass

	var funcDecl *ast.FuncDecl
	for _, decl := range pass.Files[0].Decls {
		funcDecl, _ = decl.(*ast.FuncDecl)
		if funcDecl != nil {
			break
		}
	}

	return funcDecl, pass.TypesInfo
}

func getSSAFunc(t *testing.T, filename string) *ssa.Function {
	t.Helper()

	testdata := analysistest.TestData()
	result := analysistest.Run(t, testdata, buildssa.Analyzer, filename)[0]

	ssainfo := result.Result.(*buildssa.SSA)

	return ssainfo.SrcFuncs[0]
}

func getSSADataAndFuncName(t *testing.T, filename string) (string, *SSAData) {
	t.Helper()

	testdata := analysistest.TestData()
	result := analysistest.Run(t, testdata, buildssa.Analyzer, filename)[0]

	ssainfo := result.Result.(*buildssa.SSA)

	funcName := ssainfo.SrcFuncs[0].Name()
	ssaData := GetSSAData(ssainfo)

	return funcName, ssaData
}
