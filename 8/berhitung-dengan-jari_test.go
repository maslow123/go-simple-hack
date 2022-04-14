package berhitung_dengan_jari

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBerhitungDenganJari(t *testing.T) {
	require.Equal(t, "telunjuk", berhitungDenganJari("1 10"))
	require.Equal(t, "jempol", berhitungDenganJari("3 11"))
}
