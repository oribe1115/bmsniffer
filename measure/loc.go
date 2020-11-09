package measure

import (
	"go/ast"
	"go/token"
)

// LineOfCode 対象の関数のLOCを計測する
func LineOfCode(fset *token.FileSet, funcDecl *ast.FuncDecl) int {
	startLine := fset.Position(funcDecl.Type.Func).Line
	endLint := fset.Position(funcDecl.Body.Rbrace).Line

	return endLint - startLine + 1
}
