package solutions

func isPythagoreanTriplet(a, b, c int64) bool {
	if a > b || b > c || a > c {
		return false
	}
	if a*a+b*b == c*c {
		return true
	}
	return false
}

func Problem9() int64 {
	sum := int64(1000)

	for a := int64(1); a < sum; a++ {
		for b := int64(a + 1); b < sum; b++ {
			if isPythagoreanTriplet(a, b, sum-a-b) {
				return int64(a * b * (sum - a - b))
			}
		}
	}

	return 0
}
