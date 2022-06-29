package iteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	got := Repeat("x", 5)
	expected := "xxxxx"

	assert.Equal(t, expected, got)

}
