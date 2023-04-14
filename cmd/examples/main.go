package main

import (
	"ceal/compiler"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type args struct {
	Path string
}

func run(a args) error {
	paths := []string{}

	err := filepath.WalkDir(a.Path, func(p string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			if filepath.Ext(p) == ".cpp" {
				paths = append(paths, p)
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	c := compiler.CealCompiler{
		Includes: []string{a.Path},
	}

	for _, p := range paths {
		bs, err := os.ReadFile(p)
		if err != nil {
			return err
		}

		src := string(bs)

		program := c.Compile(src)
		teal := program.String()

		out := fmt.Sprintf("%s.teal", p)
		err = os.WriteFile(out, []byte(teal), os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var a args

	flag.StringVar(&a.Path, "path", "", "source path")
	flag.Parse()

	err := run(a)
	if err != nil {
		panic(err)
	}
}
