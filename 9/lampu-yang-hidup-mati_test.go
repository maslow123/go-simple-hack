package lampu_yang_hidup_mati

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLampuYangHidupMati(t *testing.T) {
	require.Equal(t, 4, lampuYangHidupMati("4 1"))
	require.Equal(t, 4, lampuYangHidupMati("10 3"))
	require.Equal(t, 6, lampuYangHidupMati("10 4"))
}
