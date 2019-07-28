package solutions

func grid(w, h int64, cache [][]int64) int64 {
	if w <= 1 || h <= 1 {
		return h + w
	}

	if cache[w-1][h] == 0 {
		cache[w-1][h] = grid(w-1, h, cache)
	}

	if cache[w][h-1] == 0 {
		cache[w][h-1] = grid(w, h-1, cache)
	}

	return cache[w-1][h] + cache[w][h-1]
}

func Problem15() int64 {
	size := int64(20)

	cache := make([][]int64, size+1)
	for w := int64(1); w <= size; w++ {
		cache[w] = make([]int64, size+1)
	}
	return grid(size, size, cache)
}
