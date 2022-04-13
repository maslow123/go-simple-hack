package urutan_terpanjang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUrutanTerpanjang(t *testing.T) {
	require.Equal(t, 3, urutanTerpanjang("4 8 9 7 1 12 13 99 43 44"))
	require.Equal(t, 9, urutanTerpanjang("1 2 3 4 5 6 7 8 9 22 23"))
	require.Equal(t, 2, urutanTerpanjang("1 2 8 9 11 12 22 23 67 68 89 90"))
}
