package lettercase

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsCamel(t *testing.T) {
	require.True(t, IsCamel("CamelCase"))
	require.False(t, IsCamel("dromedarCase"))
	require.False(t, IsCamel("Snake_Case"))
}
