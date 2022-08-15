package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func add(val int, listNode *ListNode) *ListNode {
	var tempNode = &ListNode{
		Val:  val,
		Next: nil,
	}

	if listNode == nil {
		listNode = tempNode
	} else {
		var iterator *ListNode = listNode
		for iterator.Next != nil {
			iterator = iterator.Next
		}

		iterator.Next = tempNode
	}

	return listNode
}

func MakeSLL(arr []int) *ListNode {
	var listNode *ListNode

	for _, item := range arr {
		listNode = add(item, listNode)
	}

	return listNode
}

func length(node *ListNode) int {
	var length int = 0

	for ; node != nil; node = node.Next {
		length = length + 1
	}

	return length
}

func PrintSLL(node *ListNode) {
	fmt.Print("Single Linked List : ")

	for ; node != nil; node = node.Next {
		fmt.Print(node.Val, " -> ")
	}

	fmt.Print("\n")
}

// func main() {
// 	fmt.Print("make with array test : ", makeSLL([]int{2,3,1,4}))
// }