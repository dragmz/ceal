package ceal

import (
	"ceal/compiler"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestExamplesCompileWithoutPanic(t *testing.T) {
	paths := []string{}

	err := filepath.WalkDir("examples", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			if filepath.Ext(path) == ".cpp" {
				paths = append(paths, path)
			}
		}
		return nil
	})

	if err != nil {
		t.Error(err)
	}

	for _, path := range paths {
		bs, err := os.ReadFile(path)
		if err != nil {
			t.Error(err)
		}

		src := string(bs)

		compiler.Compile(src)
	}
}
