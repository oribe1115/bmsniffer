package measure

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func getFsetAndFuncDecl(t *testing.T, filename string) (*token.FileSet, *ast.FuncDecl, error) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, nil, err
	}

	var funcDecl *ast.FuncDecl
	for _, decl := range file.Decls {
		funcDecl, _ = decl.(*ast.FuncDecl)
		if funcDecl != nil {
			break
		}
	}

	if funcDecl == nil {
		return nil, nil, fmt.Errorf("faild to find FuncDecl")
	}

	return fset, funcDecl, nil
}

func getFsetAndFuncDeclAndInfo(t *testing.T, filename string, pkgPath string) (*token.FileSet, *ast.FuncDecl, *types.Info, error) {
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
	var conf types.Config
	_, err = conf.Check(pkgPath, fset, []*ast.File{file}, info)
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
