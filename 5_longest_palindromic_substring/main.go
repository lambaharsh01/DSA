// go run 5_longest_palindromic_substring/main.go

package main

import "fmt"

func main() {

	var str string = "sttsring ___tttaaattt___"

	fmt.Println(longestPalindrome(str))

}

func longestPalindrome(str string) string {

	var start int
	var end int

	for i := 0; i < len(str); i++ {
		charCenterLeft1, charCenterRight1 := startFromTheCenter(str, i, i)
		if (charCenterRight1 - charCenterLeft1) > (end - start) {
			end, start = charCenterRight1, charCenterLeft1
		}

		charDivideLeft1, charDivideRight1 := startFromTheCenter(str, i, i+1)
		if (charDivideRight1 - charDivideLeft1) > (end - start) {
			end, start = charDivideRight1, charDivideLeft1
		}
	}

	return str[start:end]

}

func startFromTheCenter(str string, start int, end int) (int, int) {

	for start >= 0 && end < len(str) && str[start] == str[end] {
		start--
		end++
	}

	return start + 1, end
	// here start + 1 because the left side of array SLICE EXPANSION [:] includes the index so if the start goes -1 there would be an error of out of bound

}

// BREAKTHROUGHS
// 		in SLICE EXPANSION[:] the left is inclusive index and and the right index is exclusive

// APPROACH
// 		now this problem could be approached in different ways but what i chose is a O(n²)
// 		where i think of every character and the gap between characters to be a potential center

// 		now there could be 2 thing the palindrome could start from between the 2 characters or a single character could be the center
// 		considering this to not write the code twice we can create a function which calculates the frequency of repeating substrings based on the 2 arguments from where to start the string
// 			>> same string index if the character is at the center (i, i)
// 			>> current string index ,  current string index + 1 (i, i + 1) if the character divide is in the center

// PROBLEM STATEMENT
// 		Input: s = "babad"
// 		Output: "bab"
// 		Explanation: "aba" is also a valid answer.