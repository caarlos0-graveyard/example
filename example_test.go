package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	var assert = assert.New(t)
	assert.NoError(Foo())
}
