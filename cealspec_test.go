package ceal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStackType(t *testing.T) {
	assert.Equal(t, "r32_byte_t", ReadStackType("[32]byte"))
	assert.Equal(t, "r64_byte_t", ReadStackType("[64]byte"))
	assert.Equal(t, "r_byte_t", ReadStackType("[]byte"))
}
