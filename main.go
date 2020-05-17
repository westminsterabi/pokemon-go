package main

import (
	"./pokeparser"
	"./pokeprinter"
	"./poketoken"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func createOutfile(outfile string) *os.File {
	f, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}
	return f
}

func run(outfile *os.File, args []string)  {
	goExecPath, err := exec.LookPath("go")
	progArgs := []string{goExecPath, "run", outfile.Name()}
	progArgs = append(progArgs, args...)
	cmd := exec.Cmd{
		Path: goExecPath,
		Args: progArgs,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	fileName := flag.String("filename", "", "name of the file to compile")
	flag.Parse()
	args := flag.Args()
	fset := poketoken.NewFileSet()
	exp, err := pokeparser.ParseFile(fset, *fileName, nil, 0)
	if err != nil {
		fmt.Printf("parsing failed: %s", err)
		return
	}
	fileBase := strings.TrimSuffix(*fileName, path.Ext(*fileName))
	outfileName := fmt.Sprintf("/tmp/%s.go", fileBase)
	outfile := createOutfile(outfileName)
	err = pokeprinter.Fprint(outfile, poketoken.NewFileSet(), exp)
	if err != nil {
		fmt.Printf("output failed: %s", err)
	}
	run(outfile, args)
	fmt.Printf("\n")
}