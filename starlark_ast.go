package main

import (
	"fmt"
	"os"

	"github.com/bazelbuild/buildtools/build"
	"github.com/sanity-io/litter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: starlark_ast FILENAME")
		os.Exit(1)
	}
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading file: %v\n", err)
		os.Exit(1)
	}
	file, err := build.ParseBuild(os.Args[1], contents)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing file: %v\n", err)
		os.Exit(1)
	}
	for _, stmt := range file.Stmt {
		litter.Dump(stmt)
	}

}
