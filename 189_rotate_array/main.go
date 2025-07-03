// go run 189_rotate_array/main.go

package main

// FIRST: FIND MODULUS OF K of k is more than the length of the array it has t be found out
// APPROACH 1: first make an array store all the elements after k index from the end then make and array which includes [the extracted elements including k, and then the elements before k ] hence making an array of expected output
// APPROACH 2: revers the whole array, 1,2,3,4,5,6 to 6,5,4,3,2,1 then if k is 3 then again reverse the the first 3 elements 4,5,6,3,2,1 and now reverse the elements after k th index from the beginning making it 4,5,6,1,2,3

func main(){

	nums := []int{1,2,3,4,5,6}

	rotate(nums, 2)
	
	
}

func rotate(nums []int, k int)  {
    k= k % len(nums)
    var elemLen int = len(nums) - 1
    reverse(nums, 0, elemLen)
    reverse(nums, 0, k - 1)
    reverse(nums, k, elemLen)
}

func reverse(nums []int, start int, end int){
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start ++
        end --  
    }
}

// func rotate(nums []int, k int)  {
// k= k % len(nums)
 
//  var newArray []int= nums[(len(nums)-k):];
 
//  newArray=append(newArray, nums[:len(nums)-k]...);
 
//  for i:=0; i < len(newArray); i++{
//      nums[i]=newArray[i];
//  }
// }