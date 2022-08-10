package dari_kursi_jadi_bui

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDariKursiJadiBui(t *testing.T) {
	require.Equal(t, 1, DariKursiJadiBui("BAYAM AYAM"))
	require.Equal(t, 4, DariKursiJadiBui("IKAN SAPI"))
	require.Equal(t, 4, DariKursiJadiBui("HITAM PUTIH"))
	require.Equal(t, 4, DariKursiJadiBui("KURSI BUI"))
}
