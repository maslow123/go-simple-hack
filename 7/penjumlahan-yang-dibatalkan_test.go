package penjumlahan_yang_dibatalkan

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPenjumlahanYangDibatalkan(t *testing.T) {
	require.Equal(t, 2, penjumlahanYangDibatalkan("2 2 B 10 B 7 B 5 B 11 B"))
	require.Equal(t, 11, penjumlahanYangDibatalkan("1 1 1 1 1 1 1 1 1 1 1"))
	require.Equal(t, 6, penjumlahanYangDibatalkan("10 B 2 2 B 2 2"))
}
