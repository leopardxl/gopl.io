// Package word provides utilities for word games
package word

//isPalindrome reports whether s reads the same forward and backward

func isPalindrome(s string) bool {
	length := len(s)
	for k := range s {
		if s[k] != s[length-k-1] {
			return false
		}
	}
	return true

}
