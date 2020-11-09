package measure

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// LineOfCode 対象の関数のLOCを計測する
func LineOfCode(pass *analysis.Pass, funcDecl *ast.FuncDecl) int {
	startLine := pass.Fset.Position(funcDecl.Type.Func).Line
	endLint := pass.Fset.Position(funcDecl.Body.Rbrace).Line

	return endLint - startLine + 1
}
