package palindromes

import (
	"strings"
)

func Ispalindrome(word string) bool {
	// defer timer.TimeIt("Ispalindrome")()
	n := len(word)
	for i := 0; i < n/2; i++ {
		if word[i] != word[n-1-i] {
			return false
		}
	}
	return true
}

func FindPalindromesInString(s string) []string {
	// defer timer.TimeIt("FindPalindromesInString")()

	words := strings.Split(s, " ")
	palindromes := make([]string, len(words))
	count := 0

	// Iterate over the words backwards, adding the palindromes to the slice
	// at the index count. Then increment count.
	for i := len(words) - 1; i >= 0; i-- {
		if Ispalindrome(words[i]) {
			palindromes[count] = words[i]
			count++
		}
	}

	// Get rid of the empty slots in the slice
	return palindromes[:count]
}

// In your pseudo code you didn't have a n parameter but suggest in the
// examples that one is needed. So I added it to the function.
func GetLengthOfNthPalindromeFromLast(s string, n int) int {
	// defer timer.TimeIt("GetLengthOfNthPalindromeFromLast")()
	palindromes := FindPalindromesInString(s)
	palindromeCount := len(palindromes)

	if n > palindromeCount {
		return -1
	}
	return len(palindromes[n-1])
}

// Method 1:
// 1. Gets all palindromes from the string, 0 to n.
// 2. Reverses the list of palindromes.
// 3. Returns the length of the nth palindrome from the end of the list.

// Method 2:
// 1. Get all palindromes from the string backwards to avoid the reversal.
// 2. Returns the length of the nth palindrome from the end of the list.

// Split functionality.
// First approach was to split the string on space and check each word.
// for simplicity. However, in my head this would result in two traversals.
// a more efficient approach would be to traverse the string once. Split
// on a space add to a buffer and if it is not a palindrome chuck it. However
// this approach didn't make the result any faster. Doing tests where each
// function was run 1_000_000 times, both results fluctauted around 460ms.
// Because of this I decided to keep the first approach as it is more readable

// Dynamic Allocation
// First iteration I simply used dynamic allocation make([]string, 0) and
// tried an approach where I set the length of the slice to the number of
// words because there can't be more palindromes than words. Now instead of
// appending I set the value at the index. I then slice to remove the excess
// slots. This resulted in a 510700600ns -> 420073400ns improvement. However,
// this increases memory usage (negligible).
