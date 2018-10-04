package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"été", true},
		{"palindrome", false},
	}

	for _, test := range tests {
		if got := isPalindrome(test.input); got != test.want {
			t.Errorf("isPalindrome(%q) = %v", test.input, got)
		}
	}
}
