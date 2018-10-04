package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !isPalindrome("detartrated") {
		t.Error(`isPalindrome("detartrated") = false`)
	}

	if !isPalindrome("kayak") {
		t.Error(`isPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if isPalindrome("palindrome") {
		t.Error(`isPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !isPalindrome("été") {
		t.Error(`isPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama."
	if !isPalindrome(input) {
		t.Errorf(`isPalindrome(%q) = false`, input)
	}
}
