package tanggal_yang_tak_pasti

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTanggalYangTakPasti(t *testing.T) {
	require.Equal(t, "YA", TanggalYangTakPasti("1965-3-aug"))
	require.Equal(t, "TIDAK", TanggalYangTakPasti("1965 3-aug")) // invalid format date
	require.Equal(t, "TIDAK", TanggalYangTakPasti("04/05/2010")) // ada 2 kemungkinan 4 Mei 2010 atau 5 April 2010
	require.Equal(t, "YA", TanggalYangTakPasti("2011-08-22"))
	require.Equal(t, "YA", TanggalYangTakPasti("23 mar 2000"))
	require.Equal(t, "YA", TanggalYangTakPasti("jan 3 2004"))

}
