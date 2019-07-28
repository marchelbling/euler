package solutions

func collatz(n int64) int64 {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func Problem14() int64 {
	length := map[int64]int64{
		1: 1,
		2: 2,
	}

	for i := int64(3); i < 1000000; i++ {
		// build collatz sequence
		count := int64(1)
		n := i
		for ; n != 1; n = collatz(n) {
			if n == 1 {
				length[i] = count
				break
			}
			if n < i {
				length[i] = count + length[n]
				break
			}
			count++
		}
	}

	var result, max int64
	for n, l := range length {
		if l > max {
			result = n
			max = l
		}
	}

	return result
}
