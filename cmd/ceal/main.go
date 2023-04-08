package main

import (
	"bytes"
	"ceal/compiler"
	"flag"
	"fmt"
	"os"

	"github.com/joe-p/tealfmt"
	"github.com/pkg/errors"
)

//go:generate go run ../sdk/main.go -spec ../../ceal.json --out ../../avm.hpp
//go:generate go run ../gen/main.go -spec ../../ceal.json --out ../../compiler/compiler_gen.go

type args struct {
	Path   string
	Format bool
}

func run(a args) error {
	bs, err := os.ReadFile(a.Path)
	if err != nil {
		return errors.Wrap(err, "failed to open source file")
	}

	src := string(bs)

	teal := compiler.Compile(src)

	if a.Format {
		teal = tealfmt.Format(bytes.NewBufferString(teal))
	}

	fmt.Print(teal)

	return nil
}

func main() {
	var a args
	flag.StringVar(&a.Path, "path", "", "source file path")
	flag.BoolVar(&a.Format, "format", false, "format output")
	flag.Parse()

	err := run(a)
	if err != nil {
		panic(err)
	}
}
