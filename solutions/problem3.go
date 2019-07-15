package solutions

import "github.com/marchelbling/euler/utilmath"

func Problem3() int64 {
	factors := utilmath.PrimeFactors(600851475143)
	n := len(factors)
	if n > 2 && factors[n-2] > factors[n-1] {
		return factors[n-2]
	}
	return factors[n-1]
}
