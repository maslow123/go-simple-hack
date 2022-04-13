package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceStringToSliceInt(t *testing.T) {
	require.Equal(t, []int{1, 2, 3, 4}, SliceStringToSliceInt([]string{"1", "2", "3", "4"}))
}
