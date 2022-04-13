package kata_yang_kaya

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKataYangKaya(t *testing.T) {
	require.Equal(t, 4, kataYangKaya("kata di ujung pena"))
	require.Equal(t, 5, kataYangKaya("juli di bulan juni"))
	require.Equal(t, 5, kataYangKaya("daun di atas bantal"))
}
