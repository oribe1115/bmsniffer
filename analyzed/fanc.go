package analyzed

import (
	"fmt"
	"go/ast"
)

type Func struct {
	FuncDecl   *ast.FuncDecl
	Loc        int
	Maxnesting int
	Nov        int
	Cyclo      int
}

func (fn *Func) Print() {
	fmt.Println(fn.FuncDecl.Name.String(), fn.Loc, fn.Maxnesting, fn.Nov, fn.Cyclo)
}
