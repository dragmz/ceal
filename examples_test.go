package ceal

import (
	"ceal/compiler"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
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

	for _, p := range paths {
		bs, err := os.ReadFile(p)
		if err != nil {
			t.Error(err)
		}

		src := string(bs)

		actual := compiler.Compile(src)

		tp := fmt.Sprintf("%s.teal", p)
		tbs, err := os.ReadFile(tp)
		if err != nil {
			t.Error(err)
		}

		expected := string(tbs)

		assert.Equal(t, expected, actual, fmt.Sprintf("compiled '%s' does not match '%s'", p, tp))
	}
}
