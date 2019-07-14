package solutions

func Problem1() int64 {
	var sum int64
	for i := int64(1); i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	return sum
}
