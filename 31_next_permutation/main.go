//  go run 31_next_permutation/main.go

// TBH I really did not understand the Question but still i learned what permutations are
// and a really cool trick that is in func reverse
// it was fun solving it the way chat gpt told me without the code

package main

import "fmt"

func main() {

	nums := []int{1,2,3}
	nextPermutation(nums)

	fmt.Println(nums)

}

func nextPermutation(nums []int) {

	var len int = len(nums) - 1
	var pivot int = -1

	for i := len; i > 0; i-- {

		if nums[i-1] < nums[i] {
			pivot = (i-1)
			break
		}

	}

	if pivot !=-1 {
		for i := len; i > pivot; i-- {
			if nums[i] > nums[pivot] {
				fmt.Println(i, pivot)
				nums[i] , nums[pivot] = nums[pivot], nums[i]
				break
			}
		}
	}
		
	var start int = pivot + 1

	if pivot == -1 {
		start = 0
	}
	
	reverse(nums, start, len)

}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

