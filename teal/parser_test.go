package teal

import (
	"bytes"
	"testing"

	"github.com/joe-p/tealfmt"
	"github.com/stretchr/testify/assert"
)

func testParseTo(t *testing.T, src string, expected string) {
	expected = tealfmt.Format(bytes.NewBufferString(expected))

	teal := ParseTeal(src)
	actual := tealfmt.Format(bytes.NewBufferString(teal.String()))

	assert.Equal(t, expected, actual)
}

func TestParseTo(t *testing.T) {
	type test struct {
		i string
		o string
	}

	// TODO: may actually want to encode back using the same format as the input (i.e. string instead of b64)
	tests := []test{
		{i: `byte "test"`, o: `byte b64 dGVzdA==`},
	}

	for _, test := range tests {
		testParseTo(t, test.i, test.o)
	}
}

func TestParseToSelf(t *testing.T) {
	sources := []string{
		"int 1",
		"box_create",
		`int 1
		box_create`,
		`byte b64 dGVzdA==`,
	}

	for _, source := range sources {
		testParseTo(t, source, source)
	}
}
