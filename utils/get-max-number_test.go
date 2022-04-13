package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMaxNumber(t *testing.T) {
	require.Equal(t, 10, GetMaxNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
