package lettercase

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestIsCamel(t *testing.T) {
	assert.Assert(t, IsCamel("CamelCase"))
	assert.Assert(t, !IsCamel("dromedarCase"))
	assert.Assert(t, !IsCamel("Snake_Case"))
}
