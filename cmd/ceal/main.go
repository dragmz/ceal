package main

import (
	"ceal/compiler"
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

//go:generate go run ../sdk/main.go -spec ../../ceal.json --out ../../avm.hpp
//go:generate go run ../gen/main.go -spec ../../ceal.json --out ../../compiler/compiler_gen.go

type args struct {
	Path string
}

func run(a args) error {
	bs, err := os.ReadFile(a.Path)
	if err != nil {
		return errors.Wrap(err, "failed to open source file")
	}

	src := string(bs)

	teal := compiler.Compile(src)
	fmt.Print(teal)

	return nil
}

func main() {
	var a args
	flag.StringVar(&a.Path, "path", "", "source file path")
	flag.Parse()

	err := run(a)
	if err != nil {
		panic(err)
	}
}
