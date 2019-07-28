package utilstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	for _, param := range []struct {
		value    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"ab", false},
		{"abcba", true},
		{"abcbaa", false},
	} {
		require.Equal(t, param.expected, IsPalindrome(param.value), fmt.Sprintf("failed for %q", param.value))
	}
}
