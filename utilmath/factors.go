package utilmath

import (
	"math"
	"sort"
)

func Divisors(n int64) []int64 {
	if n == 1 {
		return []int64{1}
	}

	divisors := []int64{1, n}
	for i := int64(math.Sqrt(float64(n))); i > 1; i-- {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}

	sort.Slice(divisors, func(i, j int) bool {
		return divisors[i] < divisors[j]
	})

	return divisors
}
