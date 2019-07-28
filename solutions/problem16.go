package solutions

import "math/big"

func Problem16() int64 {
	var result int64
	pow := big.NewInt(1)
	two := big.NewInt(2)

	for i := 1; i <= 1000; i++ {
		pow.Mul(pow, two)
	}

	for _, c := range pow.String() {
		result = result + int64(c-'0')
	}

	return result
}
