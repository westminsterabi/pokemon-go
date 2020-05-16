package main

import (
	"bufio"
	"flag"
	"fmt"
	"./poketoken"
	"./pokeparser"
	"./pokeprinter"
	"os"
)

func main() {
	fileName := flag.String("filename", "", "name of the file to compile")
	flag.Parse()
	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		exp, err := pokeparser.ParseExpr(line)
		if err != nil {
			fmt.Printf("parsing failed: %s", err)
			return
		}
		pokeprinter.Fprint(os.Stdout, poketoken.NewFileSet(), exp)
		fmt.Printf("\n")
		}
}
