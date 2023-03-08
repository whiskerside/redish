package redish

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHGetString(t *testing.T) {
	val := HGetString("hashset", "non-exist-key")
	require.Empty(t, val)

	vals := []string{"hk1", "v1", "hk2", "v2", "hk3", "v3"}
	HSetString("hashset", vals...)

	val = HGetString("hashset", "hk1")
	require.Equal(t, "v1", val, "return same value")
}

func TestHSetString(t *testing.T) {
	var done bool
	done = HSetString("hashset", []string{}...)
	require.False(t, done, "wrong number of arguments for 'hset' command")

	vals := []string{"k1", "v1", "k2", "v2", "k3", "v3"}
	done = HSetString("hashset", vals...)
	require.True(t, done)

	done = HSetString("hashset", []string{"k4", "v4"}...)
	require.True(t, done)
}
