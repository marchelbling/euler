package solutions

func Problem6() int64 {
	n := int64(100)
	sumOfSquares := n * (n + 1) * (2*n + 1) / 6
	squareOfSum := (n * (n + 1) / 2) * (n * (n + 1) / 2)

	return squareOfSum - sumOfSquares
}
