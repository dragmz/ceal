package main

import (
	"bytes"
	"ceal/compiler"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joe-p/tealfmt"
	"github.com/pkg/errors"
)

//go:generate go run ../sdk/main.go -spec ../../ceal.json --out ../../avm.hpp
//go:generate go run ../gen/main.go -spec ../../ceal.json --out ../../compiler/gen.go
//go:generate go run ../teal/main.go -spec ../../ceal.json --out ../../teal/gen.go

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
	c := compiler.CealCompiler{
		Includes: []string{
			filepath.Dir(a.Path),
		},
	}

	program := c.Compile(src)
	ast := program.TealAst()
	teal := ast.Teal()
	source := teal.String()

	if a.Format {
		source = tealfmt.Format(bytes.NewBufferString(source))
	}

	fmt.Print(source)

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
