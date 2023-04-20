package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
)

type args struct {
	Path string
	Out  string
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

	buf := bytes.Buffer{}

	bw := bufio.NewWriter(&buf)

	func() {
		defer bw.Flush()

		bw.WriteString(
			`package ceal
import "testing"
`)

		for _, p := range paths {
			p = filepath.ToSlash(p)
			name := filepath.Base(p[:len(p)-len(filepath.Ext(p))])
			fmt.Fprintf(bw, "func TestExample_%s(t *testing.T) {\n", name)
			fmt.Fprintf(bw, "\ttestExample(t, \"%s\")\n", p)
			bw.WriteString("}\n")
		}
	}()

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	err = os.WriteFile(a.Out, src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var a args

	flag.StringVar(&a.Path, "path", "", "source path")
	flag.StringVar(&a.Out, "out", "", "output path")
	flag.Parse()

	err := run(a)
	if err != nil {
		panic(err)
	}
}
