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


// APPROACH 
// 		1. first run 1 b search and find the matching element 
// 				treat the matched element as the center and search both left and right that way the we only get to the context and the  search it around
// 				flaws: this approach would put it to be O(n) instead of O(log n) because what if the case is [5,5,5,5,5,5,5,5] this the the first mid would match and then the linear search so called(n/2) search has to be done
// 						doing that O(log n) + O(n) would give the time complexity of O(n) to this algorithm -- failed the constraint  

// 		2. do 2 binary searches 1 finding the min border and the other max border
//				we have to find the position where nums[mid] is the target and nums[mid - 1] is not the target in finding left border for the right one nums[mid + 1] should not be the target, again the nums[mid] has to be the targe in both the cases
// 				this way there will be 2 operations O(log n) + O(log n) giving the output of O(log n)