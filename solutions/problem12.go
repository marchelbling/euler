package solutions

import (
	"github.com/marchelbling/euler/utilmath"
)

func triangleNumber(i int64) int64 {
	return int64(i * (i + 1) / 2)
}

func Problem12() int64 {
	condition := func(n int64) bool {
		return len(utilmath.Divisors(triangleNumber(n))) >= 500
	}

	for i := int64(1); ; i++ {
		if condition(i) {
			return triangleNumber(i)
		}
	}
}
