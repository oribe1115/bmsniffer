package measure

import (
	"go/ast"
)

func MaxNestingLevel(funcDecl *ast.FuncDecl) int {
	// return countNestingLevel(funcDecl.Body)
	return inspectNestLevel(funcDecl.Body)
}

func countNestingLevel(stmt ast.Stmt) int {
	if stmt == nil {
		return 0
	}

	switch stmt := stmt.(type) {
	case *ast.BlockStmt:
		var max int
		for _, childStmt := range stmt.List {
			tmp := countNestingLevel(childStmt)
			if tmp > max {
				max = tmp
			}
		}
		return max
	case *ast.IfStmt:
		max := countNestingLevel(stmt.Body)
		for childStmt := stmt.Else; childStmt != nil; {
			var tmp int
			if elseifStmt, _ := childStmt.(*ast.IfStmt); elseifStmt != nil {
				tmp = countNestingLevel(elseifStmt.Body)
				childStmt = elseifStmt.Else
			} else {
				tmp = countNestingLevel(childStmt)
				childStmt = nil
			}

			if tmp > max {
				max = tmp
			}
		}
		return max + 1
	case *ast.SwitchStmt:
		var max int
		for _, childStmt := range stmt.Body.List {
			tmp := countNestingLevel(childStmt)
			if tmp > max {
				max = tmp
			}
		}
		return max + 1
	default:
		return 0
	}
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
		}

		return true
	})

	return nestLevel
}
