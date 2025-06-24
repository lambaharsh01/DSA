// go run 19_remove_nth_node_from_end_of_list/main.go
package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){

	var rawList *ListNode = &ListNode{Val:1, Next:
				&ListNode{Val:2, Next:
					&ListNode{Val:3, Next:
						&ListNode{Val:4, Next:
							&ListNode{Val:5, Next:nil},
						},
					},
				},
			}

	listOutput := removeNthFromEnd(rawList, 2)

	for listOutput!=nil {
		fmt.Println(listOutput.Val)
		listOutput = listOutput.Next
	}


}


func removeNthFromEnd(head *ListNode, n int) *ListNode {

	var listIndex int = 0
	var list *ListNode = head

	var prevIndex int = 0
	var prev *ListNode = head

	for list !=nil {

		if listIndex > n {
			prevIndex++
			prev = prev.Next
		}

		listIndex++
		list = list.Next

	}

	if listIndex == n { // WHY THIS CONDITION: [1,2] 2
		// BREAKTHROUGH here a very specific case can occur where the prev never increases because the element to be removed is the head soo the list ends before prev list can be increased 
		head.Next = head.Next.Next
	} else if listIndex - 1 == prevIndex {
		prev.Next = nil
	}else {
		prev.Next = prev.Next.Next 
	}



	return head
}

// APPROACH
//		now according to me this could be done in multiple ways but the best approach is doing int in O(n)
//		we will have 2 variables one keeping the track of total length of the list and the other keeping the one which stays (list length) - n positions behind
// 		the listLength keeps on increasing with each iteration but the prevIndex will only increase once listLength has increased over n and remain n steps behind

// PROBLEM STATEMENT
// 		Input: head = [1,2,3,4,5], n = 2
// 		Output: [1,2,3,5]
// 		Example 2:

// 		Input: head = [1], n = 1
// 		Output: []
// 		Example 3:

// 		Input: head = [1,2], n = 1
// 		Output: [1]