package analyzed

type Pkg struct {
	Files       []*File
	FileFilters []func(*File) bool
	FuncFilters []func(*Func) bool
}

func NewPkg() *Pkg {
	return &Pkg{
		Files:       make([]*File, 0),
		FileFilters: make([]func(*File) bool, 0),
		FuncFilters: make([]func(*Func) bool, 0),
	}
}

func (pkg *Pkg) AddFile(file *File) {
	pkg.Files = append(pkg.Files, file)
}

func (pkg *Pkg) AddFileFilter(filter func(*File) bool) {
	pkg.FileFilters = append(pkg.FileFilters, filter)
}

func (pkg *Pkg) AddFuncFilter(filter func(*Func) bool) {
	pkg.FuncFilters = append(pkg.FuncFilters, filter)
}

func (pkg *Pkg) Filter() {
	newFiles := make([]*File, 0)

	for _, file := range pkg.Files {
		if pkg.filterFile(file) {
			file.FuncFilters = pkg.FuncFilters
			file.FilterFuncs()
			newFiles = append(newFiles, file)
		}
	}

	pkg.Files = newFiles
}

func (pkg *Pkg) filterFile(file *File) bool {
	for _, filter := range pkg.FileFilters {
		if !filter(file) {
			return false
		}
	}

	return true
}

func (pkg *Pkg) Print() {
	for _, file := range pkg.Files {
		file.Print()
	}
}
