package measure

import (
	"go/ast"
)

func MaxNestingLevel(funcDecl *ast.FuncDecl) int {
	return inspectNestLevel(funcDecl.Body)
}

func inspectNestLevel(n ast.Node) int {
	if n == nil {
		return 0
	}

	var nestLevel int
	ast.Inspect(n, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		switch n := n.(type) {
		case *ast.IfStmt:
			max := inspectNestLevel(n.Body)
			for child := n.Else; child != nil; {
				var tmp int
				if elseifStmt, _ := child.(*ast.IfStmt); elseifStmt != nil {
					tmp = inspectNestLevel(elseifStmt.Body)
					child = elseifStmt.Else
				} else {
					tmp = inspectNestLevel(child)
					child = nil
				}
				if tmp > max {
					max = tmp
				}
			}
			if max+1 > nestLevel {
				nestLevel = max + 1
			}
			return false
		case *ast.SwitchStmt:
			var max int
			for _, child := range n.Body.List {
				tmp := inspectNestLevel(child)
				if tmp > max {
					max = tmp
				}
			}
			if max+1 > nestLevel {
				nestLevel = max + 1
			}
			return false
		case *ast.TypeSwitchStmt:
			var max int
			for _, child := range n.Body.List {
				tmp := inspectNestLevel(child)
				if tmp > max {
					max = tmp
				}
			}
			if max+1 > nestLevel {
				nestLevel = max + 1
			}
			return false
		case *ast.SelectStmt:
			var max int
			for _, child := range n.Body.List {
				tmp := inspectNestLevel(child)
				if tmp > max {
					max = tmp
				}
			}
			if max+1 > nestLevel {
				nestLevel = max + 1
			}
			return false
		}

		return true
	})

	return nestLevel
}
