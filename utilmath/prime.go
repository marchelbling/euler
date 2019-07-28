package utilmath

import "math"

func IsPrime(value int64) bool {
	max := int64(math.Sqrt(float64(value)))

	if value == 1 {
		return false
	}
	if value == 2 {
		return true
	}
	if value%2 == 0 {
		return false
	}

	for i := int64(3); i <= max; i = i + 2 {
		if value%i == 0 {
			return false
		}
	}
	return true
}

func PrimeFactors(value int64) []int64 {
	var factors []int64

	if IsPrime(value) {
		return []int64{value}
	}

	if value%2 == 0 {
		factors = append(factors, 2)
		value = value / 2
	}

	for i := int64(3); i < value; i = i + 2 {
		if IsPrime(i) && value%i == 0 {
			factors = append(factors, i)
			value = value / i
		}
	}

	if IsPrime(value) {
		factors = append(factors, value)
	}

	return factors
}
