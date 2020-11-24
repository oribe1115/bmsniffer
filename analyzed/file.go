package analyzed

import (
	"fmt"
	"go/ast"
)

type File struct {
	Name        string
	AstFile     *ast.File
	Funcs       []*Func
	FuncFilters []func(*Func) bool
}

func NewFile(name string, astFile *ast.File) *File {
	return &File{
		Name:        name,
		AstFile:     astFile,
		Funcs:       make([]*Func, 0),
		FuncFilters: make([]func(*Func) bool, 0),
	}
}

func (file *File) AddFunc(fn *Func) {
	file.Funcs = append(file.Funcs, fn)
}

func (file *File) FilterFuncs() {
	newFuncs := make([]*Func, 0)

	for _, fn := range file.Funcs {
		if file.filterFunc(fn) {
			newFuncs = append(newFuncs, fn)
		}
	}

	file.Funcs = newFuncs
}

func (file *File) filterFunc(fn *Func) bool {
	for _, filter := range file.FuncFilters {
		if !filter(fn) {
			return false
		}
	}

	return true
}

func (file *File) Print() {
	fmt.Printf("--- %s: LOC-MAXNESTING-NOAV-CYCLO\n", file.Name)

	for _, fn := range file.Funcs {
		fn.Print()
	}
}
