package utilmath

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDivisors(t *testing.T) {
	for _, param := range []struct {
		value    int64
		expected []int64
	}{
		{1, []int64{1}},
		{3, []int64{1, 3}},
		{4, []int64{1, 2, 4}},
		{10, []int64{1, 2, 5, 10}},
		{15, []int64{1, 3, 5, 15}},
		{21, []int64{1, 3, 7, 21}},
		{28, []int64{1, 2, 4, 7, 14, 28}},
	} {
		require.Equal(t, param.expected, Divisors(param.value))
	}
}
