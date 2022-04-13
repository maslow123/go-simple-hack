package percakapan_rahasia

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPercakapanRahasia(t *testing.T) {
	require.Equal(t, "melati dari jayagiri", percakapanRahasia("italem irad irigayaj"))
	require.Equal(t, "badai pasti berlalu", percakapanRahasia("iadab itsap ulalreb"))
	require.Equal(t, "bulan tertusuk ilalang", percakapanRahasia("nalub kusutret gnalali"))
	require.Equal(t, "aku sayang ibu dan bapak", percakapanRahasia("uka gnayas ubi nad kapab"))
}
