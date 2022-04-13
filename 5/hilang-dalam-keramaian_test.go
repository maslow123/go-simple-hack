package hilang_dalam_keramaian

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHilangDalamKeramaian(t *testing.T) {
	unitDigitMissingNumber := hilangDalamKeramaian("012346789")
	// unitDigitMissingNumber := hilangDalamKeramaian("023456789")
	tensDigitMissingNumber := hilangDalamKeramaian("23242526272830")
	// tensDigitMissingNumber := hilangDalamKeramaian("2123242526272830")
	hundredDigitMissingNumber := hilangDalamKeramaian("101102103104105106107108109111112")
	// hundredDigitMissingNumber := hilangDalamKeramaian("101103104105106107108109110111112")
	invalidNumber := hilangDalamKeramaian("123910257123091238")

	require.Equal(t, 5, unitDigitMissingNumber)
	// require.Equal(t, 1, unitDigitMissingNumber)
	require.Equal(t, 29, tensDigitMissingNumber)
	// require.Equal(t, 22, tensDigitMissingNumber)
	require.Equal(t, 110, hundredDigitMissingNumber)
	// require.Equal(t, 102, hundredDigitMissingNumber)
	require.Equal(t, 0, invalidNumber)
}
