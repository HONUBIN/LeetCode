package main

import "fmt"

func main() {
	// node_5 := ListNode{
	// 	Val: 0,
	// }
	// node_4 := ListNode{
	// 	Val:  4,
	// 	Next: &node_5,
	// }
	node_3 := ListNode{
		Val: 4,
		// Next: &node_4,
	}
	node_2 := ListNode{
		Val:  2,
		Next: &node_3,
	}
	node_1 := ListNode{
		Val:  3,
		Next: &node_2,
	}

	fmt.Println(insertionSortList(&node_1))
}
