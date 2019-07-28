package solutions

import (
	"fmt"

	"github.com/marchelbling/euler/utilstring"
)

func Problem4() int64 {
	var result int64

	for i := int64(100); i < 1000; i++ {
		for j := int64(100); j < 1000; j++ {
			product := i * j
			if utilstring.IsPalindrome(fmt.Sprintf("%d", product)) && product > result {
				result = product
			}
		}
	}
	return result
}
