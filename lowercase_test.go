package pggoquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	result := Lowercase("Hello_WORLD")
	assert.Equal(t, "hello_world", result, "The Split function should convert input to lowercase")
}
