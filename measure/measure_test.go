package measure

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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
