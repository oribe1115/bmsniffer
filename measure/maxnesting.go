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

			nestLevel = getMax(nestLevel, max+1)

			return false
		case *ast.SwitchStmt:
			max := inlChildren(n.Body.List)
			nestLevel = getMax(nestLevel, max+1)
			return false
		case *ast.TypeSwitchStmt:
			max := inlChildren(n.Body.List)
			nestLevel = getMax(nestLevel, max+1)
			return false
		case *ast.SelectStmt:
			max := inlChildren(n.Body.List)
			nestLevel = getMax(nestLevel, max+1)
			return false
		}

		return true
	})

	return nestLevel
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func inlChildren(children []ast.Stmt) int {
	var max int
	for _, child := range children {
		tmp := inspectNestLevel(child)
		if tmp > max {
			max = tmp
		}
	}
	return max
}
