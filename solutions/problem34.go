package solutions

import (
	"fmt"
)

func isNumEqualToDigitFactorialSum(n int64) bool {
	factorials := map[int64]int64{
		0: 1,
		1: 1,
		2: 2,
		3: 6,
		4: 24,
		5: 120,
		6: 720,
		7: 5040,
		8: 40320,
		9: 362880,
	}

	value := n
	sum := int64(0)
	for n > 0 {
		digit := n % 10
		n = n / 10
		sum += factorials[digit]
	}
	return sum == value
}

func Problem34() int64 {
	sum := int64(0)
	for i := int64(1); i < 100_000_000; i++ {
		if isNumEqualToDigitFactorialSum(i) {
			fmt.Printf("adding %d to the list\n", i)
			sum += int64(i)
		}
	}

	return sum
}
