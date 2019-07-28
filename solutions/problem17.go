package solutions

import (
	"log"
	"strings"
)

func toWord(n int) string {
	units := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	tens := map[int]string{
		1: "ten",
		2: "twenty",
		3: "thirty",
		4: "forty",
		5: "fifty",
		6: "sixty",
		7: "seventy",
		8: "eighty",
		9: "ninety",
	}

	exceptions := map[int]string{
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}

	result := ""
	if n >= 100 {
		h := n / 100
		n = n % 100
		result = result + " " + units[h] + " hundred"

		if n > 0 {
			result = result + " and"
		}
	}

	if n >= 10 {
		if _, ok := exceptions[n]; ok {
			result = result + " " + exceptions[n]
			n = 0
		} else {
			t := n / 10
			n = n % 10
			result = result + " " + tens[t]
		}
	}

	if n > 0 {
		result = result + " " + units[n]
	}

	log.Printf(result)
	return result
}

func lengthWithoutSpace(s string) int64 {
	return int64(len(s) - strings.Count(s, " "))
}

func Problem17() int64 {
	result := lengthWithoutSpace("one thousand")

	for i := 1; i < 1000; i++ {
		result += lengthWithoutSpace(toWord(i))
	}

	return result
}
