// go run 24_swap_nodes_in_pairs/main.go

package main

import "fmt"

// head = [1, 2, 3]
// Output: [2, 1, 3]

// head = [1, 2, 3, 4]
// Output: [2, 1, 4, 3]

// APPROACH
// we will just have a pointer referring to the list before the pair when switching 3,4 to 4,3 well have a pointer pointing to 1(after having been swapped with 2)
// the pointer will jump by 2 steps
// the pointer.Next(first item) and current(second item)
// first the first items value will be second items value (first items value being stored before being swapped)
// then the second item will be assigned the value of the first item



type ListNode struct {
	Val int
	Next *ListNode
}

func main(){

	head := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				// Val: 3, Next: &ListNode{
					Val: 4, Next: nil,
				},
			// },
		},
	}

	fmt.Println("Running")
	
	newHead := swapPairs(head)
	
	for newHead !=nil {

		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}

}
func swapPairs(head *ListNode) *ListNode {

	var prev *ListNode
	list := head

	var count int

	for list != nil {

		if prev != nil && count == 1 {

			first := prev.Next // first having the access to the first node
			first.Next = list.Next // reassigning the Next of first node(which is the second node itself) to the Next of the second's Next node to bypass the second node  
			list.Next = first // then assign the first node as the next of the second node
			prev.Next = list // once done re assign the second node as the child of the parent node of first node 

			prev = first // now if only works at every second node soo now as the swap happened now the first node is the second node so we swap so that the list = list.Next can sustain 
			list = first // now we assign the second node as the parent for the next swap
			count = 0
		}else if count == 1 {

			first := head
			first.Next = list.Next
			list.Next = first
			head = list
			
			prev = first
			list = first
			count = 0
		}else {
			count++
		}

		list = list.Next
	}
	return head
}


// func swapPairs(head *ListNode) *ListNode {

//     var list *ListNode = head
// 	var prev2 *ListNode = nil
//     var count int

// 	var i int
//     for list!=nil {

// 		// fmt.Println(count, list.Val)

// 		if i > 5 {
// 			return nil
// 		}

// 		i ++

// 		if prev2 != nil && count == 1 {

// 			fmt.Println(prev2.Val, list.Val)

// 			var prev *ListNode = prev2.Next
// 			prev.Next = list.Next
// 			prev2.Next = list
// 			list.Next = prev

// 			prev2 = list
// 			count = 0

// 		} else if count == 1 {

// 			var prev *ListNode = head
// 			prev.Next = list.Next
// 			head = list
// 			list.Next = prev

// 			prev2 = prev
// 			count = 0

// 			list = prev

// 		}else {
// 			// fmt.Println(count, list.Val)
// 			count++
// 		}

// 		// fmt.Println(list, list.Next, head)

// 		list = list.Next
//     }

// 	return head
// }