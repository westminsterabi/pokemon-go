package main

import (
	"./pokeparser"
	"./pokeprinter"
	"./poketoken"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("filename", "", "name of the file to compile")
	flag.Parse()
	fset := poketoken.NewFileSet()
	exp, err := pokeparser.ParseFile(fset, *fileName, nil, 0)
	if err != nil {
		fmt.Printf("parsing failed: %s", err)
		return
	}
	err = pokeprinter.Fprint(os.Stdout, poketoken.NewFileSet(), exp)
	if err != nil {
		fmt.Printf("output failed: %s", err)
	}
	fmt.Printf("\n")
}