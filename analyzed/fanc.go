package analyzed

import (
	"fmt"
	"go/ast"
)

type Func struct {
	FuncDecl   *ast.FuncDecl
	Loc        int
	Maxnesting int
	Noav       int
	Cyclo      int
}

func (fn *Func) Print() {
	fmt.Printf("  %s: %d-%d-%d-%d\n", fn.FuncDecl.Name.String(), fn.Loc, fn.Maxnesting, fn.Noav, fn.Cyclo)
}
