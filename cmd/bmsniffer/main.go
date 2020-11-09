package main

import (
	"github.com/oribe1115/bmsniffer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(bmsniffer.Analyzer) }

