package main

import (
	"fmt"
	"io"
	"os"
	"log"
)

func dirTree(out io.Reader, path string, printFiles bool) error {
	file, err := os.Open("testdata")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file.Readdir(0))

	return err
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
