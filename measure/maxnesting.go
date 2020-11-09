package measure

import (
	"go/ast"
)

func MaxNestingLevel(funcDecl *ast.FuncDecl) int {
	return countNestingLevel(funcDecl.Body)
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
