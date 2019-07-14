package utilmath

import "context"

func Fibonacci(ctx context.Context) <-chan int64 {
	sequence := make(chan int64)

	go func() {
		v1 := int64(1)
		v2 := int64(2)
		sequence <- v1
		sequence <- v2
		for {
			select {
			case <-ctx.Done():
				close(sequence)
				return
			default:
				sequence <- v1 + v2
				v1, v2 = v2, v1+v2
			}
		}
	}()

	return sequence
}
