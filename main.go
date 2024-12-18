package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/thiago-figueredo/normpath/pathlib"
)

func usage() {
	fmt.Println("Usage: main.go <dirpath>")
}

func validateArgs(args []string) *os.File {
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}

	file, err := os.Open(args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return file
}

func main() {
	file := validateArgs(os.Args)
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	absPath, err := filepath.Abs(cwd + string(filepath.Separator) + file.Name())

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	normpath := pathlib.Normpath(absPath)

	if filepath.Base(normpath) != filepath.Base(cwd) {
		normpath = cwd
	}

	dir, err := os.ReadDir(normpath)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Entries of `%s`: \n\n", normpath)

	for _, entry := range dir {
		info, err := entry.Info()

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("%s %s %s\n", entry.Type().String(), info.ModTime().Format("Jan 2 15:04"), entry.Name())
	}
}
