package a

import (
	"testing"

	"github.com/stretchr/testify/assert" // want `use test assertions from "gotest.tools/v3" instead of "github.com/stretchr/testify" \(see https://pkg.go.dev/gotest.tools/v3/assert\)`
)

func IllegalImport(t *testing.T) {
	assert.Equal(t, 1, 1)
}
