// go run 75_sort_colors/main.go
package main

import "fmt"

// GIVE ARRAY OF 0s,1s, and 2s
// THE TASK IS TO SORT THEM IN PLACE
// BEST CASE SINGLE PASS O(n) time and O(1) space

func main(){
	var nums []int = []int{1,2,1,2,1,2,1,2,0,0,0,0,0,2,2,2,2}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int){

	// take 3 pointers one to keep track of the beginning the current index(if not replacing 2) and high to keep tack of the end of the array
	
	var low, mid int
	var high int = len(nums)-1

	for mid <= high { // until mid is more than high (the mid has not reached 2)

		if nums[mid] == 0 { // if 0 is found the mid is replaced with the start and the middle value
			nums[mid], nums[low] = nums[low], nums[mid]
			low++
			mid++
		} else if nums[mid] == 1 { // if mid is 1 nothing happens ut mid increased
			mid++
		} else { // now if the mid is 2 the high is replaced with mid but the mid stays where it is checking in the next iteration when if the swaped element is 
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}


	// CASES WHICH CONFUSED ME {
	// 	WHAT IF mid val is 2 and the high val is 2 so will they be stuck in a loop of ever changing themselves
	// 	and the answer is NO because once replace the mid stays 2 and high gets -- mean h = h-1 so it will be effective on the index previous to it now and the swap will take place again and again until high reaches mid and the loop stops  
	// }

}