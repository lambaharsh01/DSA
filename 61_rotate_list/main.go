// go run 61_rotate_list/main.go
package main

import "fmt"

// Given the head of a linked list, rotate the list to the right by k places.
// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]

// BRUTE FORCE
//  take 3 pointers
// 	1: for the list to iterate
// 	2: for the pointer pointing to the previous which start increasing after k indexes are passes
// 	3: for keeping the track of the last list

// 	and an index variable of type int which keeps the track of the number of elements // similar to len for list

// once done we hae the element who's next will be the first and

type ListNode struct {
    Val int
    Next *ListNode
}

func main(){

	var rawList *ListNode = &ListNode{Val:1, Next:
			&ListNode{Val:2, Next: nil,
			// 	&ListNode{Val:2, Next: nil,
			// 		// &ListNode{Val:4, Next:
			// 		// 	&ListNode{Val:5, Next:nil},
			// 		// },
			// 	},
			},
		}

	head := rotateRight(rawList, 2)

	for head!=nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil {
		return nil
	}

    var prev *ListNode = head

    var list *ListNode = head
    var index int = 0

	for list.Next!=nil {
		if index > k {
			prev = prev.Next
		}
		list = list.Next
		index++
	}

	if index > k {
		prev = prev.Next
	}
	index++


	if index == k {
		return head
	}
	
	if k > index {
		k = k % index
		return rotateRight(head, k)
	}
	
	list.Next = head  
	head = prev.Next
	prev.Next = nil

	return head
    
}