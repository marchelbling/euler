package solutions

import "github.com/marchelbling/euler/utilmath"

func Problem10() int64 {
	max := int64(2000000)
	sum := int64(2)

	for i := int64(3); i < max; i = i + 2 {
		if utilmath.IsPrime(i) {
			sum = sum + i
		}
	}

	return sum
}
