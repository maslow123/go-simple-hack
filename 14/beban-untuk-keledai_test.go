package beban_untuk_keledai

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBebanUntukKeledai(t *testing.T) {
	require.Equal(t, 0, BebanUntukKeledai("10 10 20 15 5 5 10 20 20 20 15 10"))
	require.Equal(t, 10, BebanUntukKeledai("40 80 10 10 10"))
	require.Equal(t, 5, BebanUntukKeledai("5 10"))
	require.Equal(t, 0, BebanUntukKeledai("2 2 4 4 6 6 8 8 2 2"))

	invalidData := dummyData(0)
	require.Equal(t, -1, BebanUntukKeledai(invalidData)) // invalid input the data

	invalidData = dummyData(101)
	require.Equal(t, -1, BebanUntukKeledai(invalidData)) // invalid input the data
}

func dummyData(number int) string {
	data := ""
	for i := 1; i <= number; i++ {
		data = fmt.Sprintf("%s %d", data, i)
	}

	return data
}
