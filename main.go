package main

import (
	"bufio"
	"flag"
	"fmt"
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
		fmt.Println(fileScanner.Text())
	}
}
