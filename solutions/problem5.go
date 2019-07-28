package solutions

func evenlyDivisible(value, n int64) bool {
	for i := int64(2); i <= n; i++ {
		if value%i != 0 {
			return false
		}
	}
	return true
}

func Problem5() int64 {
	var result int64

	// first value to test is the product of all primes below 20
	result = 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19

	for {
		if evenlyDivisible(result, 20) {
			return result
		}
		result++
	}
}
