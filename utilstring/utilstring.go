package utilstring

func IsPalindrome(v string) bool {
	first := 0
	last := len(v) - 1
	for first < last {
		if v[first] != v[last] {
			return false
		}
		first++
		last--
	}
	return true
}
