package main

import "fmt"

func main() {
	var arr []int = []int{1, 3, 5, 7, 9, 11, 13}
	var target int = 7
	fmt.Println(BinarySearch(arr, target))
}

func BinarySearch(arr []int, target int) int {

	var left int = 0
	var right int = len(arr) - 1

	for left <= right {
		var mid int = (left + right) / 2

		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1

}
