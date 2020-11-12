package measure

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
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

func getFsetAndFuncDeclAndInfo(t *testing.T, filename string) (*token.FileSet, *ast.FuncDecl, *types.Info, error) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, nil, nil, err
	}

	info := &types.Info{
		Types:  make(map[ast.Expr]types.TypeAndValue),
		Defs:   make(map[*ast.Ident]types.Object),
		Uses:   make(map[*ast.Ident]types.Object),
		Scopes: make(map[ast.Node]*types.Scope),
	}
	conf := &types.Config{Importer: importer.Default()}
	_, err = conf.Check("", fset, []*ast.File{file}, info)
	if err != nil {
		return nil, nil, nil, err
	}

	var funcDecl *ast.FuncDecl
	for _, decl := range file.Decls {
		funcDecl, _ = decl.(*ast.FuncDecl)
		if funcDecl != nil {
			break
		}
	}

	if funcDecl == nil {
		return nil, nil, nil, fmt.Errorf("faild to find FuncDecl")
	}

	return fset, funcDecl, info, nil
}

func getSSAFunc(t *testing.T, filename string) *ssa.Function {
	t.Helper()

	testdata := analysistest.TestData()
	result := analysistest.Run(t, testdata, buildssa.Analyzer, filename)[0]

	ssainfo := result.Result.(*buildssa.SSA)

	return ssainfo.SrcFuncs[0]
}
