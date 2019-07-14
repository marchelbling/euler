package solutions

import (
	"context"

	"github.com/marchelbling/euler/utilmath"
)

func Problem2() int64 {
	var sum int64
	ctx, cancel := context.WithCancel(context.Background())

	for v := range utilmath.Fibonacci(ctx) {
		if v >= 4000000 {
			cancel()
		}
		if v%2 == 0 {
			sum = sum + v
		}
	}
	return sum
}
