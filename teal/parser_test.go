package teal

import (
	"bytes"
	"testing"

	"github.com/joe-p/tealfmt"
	"github.com/stretchr/testify/assert"
)

func testParseToSelf(t *testing.T, src string) {
	expected := tealfmt.Format(bytes.NewBufferString(src))

	teal := ParseTeal(src)
	actual := tealfmt.Format(bytes.NewBufferString(teal.String()))

	assert.Equal(t, expected, actual)
}

func TestParseToSelf(t *testing.T) {
	sources := []string{
		"int 1",
		"box_create",
		`int 1
		box_create`,
	}

	for _, source := range sources {
		testParseToSelf(t, source)
	}
}
