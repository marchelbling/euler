package solutions

import (
	"github.com/marchelbling/euler/utilmath"
)

func Problem7() int64 {
	// start at 3 to ignore 2 as it is the only even prime.
	count := int64(1) // first prime is 2
	i := int64(3)

	for {
		if utilmath.IsPrime(i) {
			count++
		}

		if count == 10001 {
			return i
		}

		i++
	}
}
