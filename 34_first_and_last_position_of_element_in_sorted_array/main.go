// go run 34_first_and_last_position_of_element_in_sorted_array/main.go
package main

import "fmt"

func main() {

	var nums []int = []int{1,2,3,4,4,4,6,7,7,7,8} 


	fmt.Println(searchRange(nums, 4))



	
}

func searchRange(nums []int, target int) []int {

	var min, max int = -1, -1
	var left, right int = 0, len(nums) - 1

	for left <= right {
		var mid int = (left + right) / 2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}else {
			if mid == 0 {
				min = mid
				break
			}else if nums[mid - 1] != target {
				min = mid
				break
			}else {
				right = mid - 1
			}
		}
	}


	left = 0
	right = len(nums) - 1

	for left <= right {
		var mid int = (left + right) / 2

		if nums[mid] < target {
			left = mid + 1 
		}else if nums[mid] > target {
			right = mid - 1
		}else {
			if mid == len(nums) - 1 {
				max = mid
				break
			}else if nums[mid + 1] != target {
				max = mid
				break
			}else {
				left = mid + 1
			}
		}
	}



	return []int{min, max}
}