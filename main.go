package main

import (
	"fmt"
	"ondo/palindromes"
)

func main() {
	const str string = "memory ordering in programs is often a complex area to deal with. malayalam and madam are palindromes. refer c++ books"

	// defer timer.TimeIt("main")()
	res := palindromes.GetLengthOfNthPalindromeFromLast(str, 3)
	fmt.Println(res)

}
