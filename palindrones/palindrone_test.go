package palindromes

import (
	"testing"
)

func TestGetLengthOfNthPalindromeFromLast(t *testing.T) {
	// Test case
	var want int
	var got int
	str := "memory ordering in programs is often a complex area to deal with. malayalam and madam are palindromes. refer c++ books"

	// Test case 1
	want = 9
	got = GetLengthOfNthPalindromeFromLast(str, 3)
	if got != want {
		t.Errorf("Test case 1 failed")
	}

	// Test case 2
	want = 5
	got = GetLengthOfNthPalindromeFromLast(str, 1)
	if got != want {
		t.Errorf("Test case 2 failed")
	}

	// Test case 3
	want = -1
	got = GetLengthOfNthPalindromeFromLast(str, 5)
	if got != want {
		t.Errorf("Test case 3 failed")
	}
}

func TestIspalindrome(t *testing.T) {
	// Test case
	var want bool
	var got bool

	// Test case 1
	want = true
	got = Ispalindrome("madam")
	if got != want {
		t.Errorf("Test case 1 failed")
	}

	// Test case 2
	want = false
	got = Ispalindrome("madame")
	if got != want {
		t.Errorf("Test case 2 failed")
	}

	// Test case 3
	want = true
	got = Ispalindrome("malayalam")
	if got != want {
		t.Errorf("Test case 3 failed")
	}
}

func TestFindPalindromesInString(t *testing.T) {
	// Test case
	var want []string
	var got []string
	str := "memory ordering in programs is often a complex area to deal with. malayalam and madam are palindromes. refer c++ books"

	// Apparently palindromes also include single characters...
	want = []string{"refer", "madam", "malayalam", "a"}

	got = FindPalindromesInString(str)
	if len(got) != len(want) {
		t.Errorf("Test case 1 failed, %s", got)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Test case 1 failed, want: %s, got: %s", want[i], got[i])
		}
	}
}
