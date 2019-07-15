package utilmath

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPrime(t *testing.T) {
	for _, param := range []struct {
		value    int64
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{6, false},
		{17, true},
		{600851475143, false},
	} {
		require.Equal(t, param.expected, IsPrime(param.value))
	}
}

func TestPrimeFactors(t *testing.T) {
	for _, param := range []struct {
		value    int64
		expected []int64
	}{
		{3, []int64{3}},
		{6, []int64{2, 3}},
		{1155, []int64{3, 5, 7, 11}},
		{13195, []int64{5, 7, 13, 29}},
	} {
		require.Equal(t, param.expected, PrimeFactors(param.value))
	}
}
