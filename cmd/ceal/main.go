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

type args struct {
	Path    string
	Format  bool
	Include string
}

func run(a args) error {
	bs, err := os.ReadFile(a.Path)
	if err != nil {
		return errors.Wrap(err, "failed to open source file")
	}

	src := string(bs)
	c := compiler.NewCompiler(compiler.CealCompilerConfig{
		Includes: []string{
			".",
			filepath.Dir(a.Path),
			filepath.Dir(a.Include),
		},
	})

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
	flag.StringVar(&a.Include, "include", "", "include path")
	flag.Parse()

	err := run(a)
	if err != nil {
		panic(err)
	}
}
