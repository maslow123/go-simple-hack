package menutup_yang_terbuka

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMenutupYangTerbuka(t *testing.T) {
	require.Equal(t, true, menutupYangTerbuka("hkd ( ffg { jknn } sdfhngjtttg )"))
	require.Equal(t, true, menutupYangTerbuka("h ((( sss []( c )))) xyz"))
	require.Equal(t, false, menutupYangTerbuka("hhiwwdf [ dfgg ) bqwci } dff"))
}
