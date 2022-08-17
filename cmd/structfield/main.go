package main

import (
	"github.com/0xjac/structfield"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(structfield.Analyzer)
}
