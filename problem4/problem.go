package problem4

import "fmt"

func IsPalindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func Run() int {
	var maxPalindromeNumber int = -1
	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			if IsPalindrome(fmt.Sprintf("%d", i*j)) {
				palindromeNumber := i * j
				if palindromeNumber > maxPalindromeNumber {
					maxPalindromeNumber = palindromeNumber
				}
			}
		}
	}
	return maxPalindromeNumber
}
