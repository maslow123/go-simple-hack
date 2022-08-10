package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
	require.Equal(t, []int{1, 2, 3}, BubbleSort([]int{3, 1, 2}))
}
