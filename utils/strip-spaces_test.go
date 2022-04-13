package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripChar(t *testing.T) {
	require.Equal(t, "omama", StripSpaces("o m a m a"))
}
