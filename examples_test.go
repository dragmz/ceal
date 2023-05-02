package ceal

import (
	"ceal/compiler"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testExample(t *testing.T, p string) {
	c := compiler.CealCompiler{
		Includes: []string{"examples", "."},
	}

	bs, err := os.ReadFile(p)
	if err != nil {
		t.Error(err)
	}

	src := string(bs)

	program := c.Compile(src)
	ast := program.TealAst()
	teal := ast.Teal()
	actual := teal.String()

	tp := fmt.Sprintf("%s.teal", p)
	tbs, err := os.ReadFile(tp)
	if err != nil {
		t.Error(err)
	}

	expected := string(tbs)

	assert.Equal(t, expected, actual, fmt.Sprintf("compiled '%s' does not match '%s'", p, tp))
}
