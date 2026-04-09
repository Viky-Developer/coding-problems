package palindrome

import "strings"

// PalindromeBruteForce checks whether a given string is a palindrome
// by comparing characters from both ends towards the center.
//
// The input string is first converted to lowercase to ensure
// case-insensitive comparison.
//
// Example:
//
//	Input:  "Madam"
//	Output: true
//
//	Input:  "Hello"
//	Output: false
//
// Time Complexity
// O(n)
// Iterates through half the string → still linear
// Space Complexity
// O(n)
// strings.ToLower() creates a new string
func PalindromeBruteForce(s string) bool {

	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {

		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

// Palindrome checks whether a given integer is a palindrome.
//
// A number is considered a palindrome if it reads the same
// forward and backward.
//
// Negative numbers and numbers ending with 0 (except 0 itself)
// are not palindromes.
//
// Example:
//
//	Input:  121
//	Output: true
//
//	Input:  -121
//	Output: false
//
// Time Complexity
// O(log10 n)
// Number of digits in the integer
// Space Complexity
// O(1)
// No extra space used
func Palindrome(value int) bool {

	if value < 0 || (value%10 == 0 && value != 0) {
		return false
	}

	original, reversed := value, 0
	for value != 0 {

		reversed = reversed*10 + value%10

		value = value / 10
	}

	return original == reversed
}

// PalindromeString checks whether a string is a palindrome
// using rune slicing to properly support Unicode characters.
//
// It compares characters from both ends moving toward the center.
//
// Example:
//
//	Input:  "madam"
//	Output: true
//
//	Input:  "hello"
//	Output: false
//
// Time Complexity
// O(n)
// Space Complexity
// O(n)
// Creating []rune slice
func PalindromeString(value string) bool {

	runes := []rune(value)

	start, end := 0, len(runes)-1

	for start < end {

		if runes[start] != runes[end] {
			return false
		}

		start++
		end--
	}
	return true
}
