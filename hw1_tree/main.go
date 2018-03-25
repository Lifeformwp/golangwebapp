package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"io/ioutil"
)

func dir(file []os.FileInfo, path string) error {
	for _, val := range file {
		if val.IsDir() != true {
			fmt.Println("├───" + val.Name())
		} else {
			fmt.Println("├───" + val.Name())
			file, _ := ioutil.ReadDir(path + val.Name())
			dir(file, path + val.Name() + "/")
		}
	}

	return nil
}

func dirTree(out io.Reader, path string, printFiles bool) error {
	file, err := ioutil.ReadDir("testdata")
	if err != nil {
		log.Fatal(err)
	}

    dir(file, "testdata/")

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
