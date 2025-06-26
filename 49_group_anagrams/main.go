// go run 49_group_anagrams/main.go
package main

import "fmt"

// BREAKTHROUGH
// HOW TO GET THE INDEX OF AN ALPHABET(only for lower case letters)
// var char rune = 'x'
// int(char - 'a') // would give 23 (a=0 AND z=25)

// BREAKTHROUGH 2
// fmt.Sprintf("%v", arg) convert any value into a string using its default format.
// name := "Harsh" fmt.Sprintf("%v", name)  // "Harsh"
// age := 23 fmt.Sprintf("%v", age) // "23"
// arr := []int{1, 2, 3} fmt.Sprintf("%v", arr) // "[1 2 3]"
// Use %#v if you want the Go syntax representation: fmt.Sprintf("%#v", arr)  // returns "[]int{1, 2, 3}"


func main(){
	var strs []string = []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams(strs))
}



// APPROACH 2. COUNT VECTOR HASHING METHOD
// this approach is like having an array of length 26 init at all 0 values and then the character length ++ according to the letters index in the alphabet array
func groupAnagrams(strs []string) [][]string {

	maps := make(map[string][]string, len(strs))

	for _, str := range strs {

		var charSet [26]int

		for _, char := range str {
			charSet[int(char - 'a')] += 1
		} 

		stringed := fmt.Sprintf("%v", charSet)

		if _, ok := maps[stringed]; ok {
			maps[stringed] = append(maps[stringed], str)
		}else {
			maps[stringed] = []string{str}
		}
		
	}

	var arrGroups [][]string

	for _, groups := range maps {
		arrGroups = append(arrGroups, groups)
	}
	
	return arrGroups
}

// APPROACH 1. SORT AND HASH METHOD

// func groupAnagrams(strs []string) [][]string {

//     maps := make(map[string][]string, len(strs))

// 	for _, str := range strs{
// 		var sortedStr string = string( quickSort([]rune(str)) )
// 		if _, ok := maps[sortedStr]; ok {
// 			maps[sortedStr] = append(maps[sortedStr], str)
// 		} else {
// 			maps[sortedStr] = []string{str}
// 		}
// 	}

// 	var returnStrs [][]string
// 	for _, strArr := range maps {
// 		returnStrs = append(returnStrs, strArr)
// 	}

// 	return returnStrs    
// }

// func quickSort(runes []rune) []rune {
// 	var mid int = len(runes) /2

// 	if mid == 0 {
// 		return runes
// 	}

// 	var pivot rune = runes[mid]

// 	var left []rune
// 	var right []rune

// 	for i, char := range runes {

// 		if i == mid {
// 			continue
// 		}

// 		if char > pivot {
// 			right = append(right, char)
// 		} else {
// 			left = append(left, char)
// 		}
// 	}

// 	return append( append(quickSort(left), pivot), quickSort(right)...)
// }