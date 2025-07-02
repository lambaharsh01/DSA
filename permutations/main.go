// go run permutations/main.go

package main

import "fmt"

func main() {

	var nums []int = []int{1, 2, 3}
	permutate(nums)

	fmt.Println(nums)

}

func permutate(nums []int) [][]int {

	var result [][]int
	var backTrack func(int)

	backTrack = func(start int) {
		if start == len(nums) {
			
			dup := make([]int, len(nums))
			copy(dup, nums)
			result = append(result, dup)
			return
		}

		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backTrack(start + 1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}

	backTrack(0)

	return result
}