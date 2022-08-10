package jumlah_jam_kerja

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJumlahJamKerja(t *testing.T) {
	require.Equal(t, 0, JumlahJamKerja("1 4 4")) // invalid input
	require.Equal(t, 0, JumlahJamKerja("aaa 0")) // invalid clock in
	require.Equal(t, 0, JumlahJamKerja("-1 0"))  // invalid clock in
	require.Equal(t, 0, JumlahJamKerja("0 aa"))  // invalid clock out
	require.Equal(t, 0, JumlahJamKerja("0 -1"))  // invalid clock out
	require.Equal(t, 0, JumlahJamKerja("10 5"))  // clock in greather than clock out

	require.Equal(t, 11, JumlahJamKerja("10 20"))
	require.Equal(t, 3, JumlahJamKerja("7 10"))
	require.Equal(t, 9, JumlahJamKerja("1 10"))

}
