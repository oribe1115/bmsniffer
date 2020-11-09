package measure

import (
	"go/ast"
	"go/types"
)

var objExists = struct{}{}

func NumberOfAccessedVariables(funcDecl *ast.FuncDecl, info *types.Info) int {
	usedVarMap := map[types.Object]struct{}{}

	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		switch n := n.(type) {
		case *ast.Ident:
			obj := info.ObjectOf(n)
			usedVarMap[obj] = objExists
		}

		return true
	})

	return len(usedVarMap)
}
