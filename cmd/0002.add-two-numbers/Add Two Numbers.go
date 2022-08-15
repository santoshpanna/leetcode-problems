package main

import utils "github.com/santoshpanna/leetcode-problems/pkg/utils"

func add(val int, node *utils.ListNode) *utils.ListNode {
	var tempNode = &utils.ListNode{
		Val:  val,
		Next: nil,
	}

	if node == nil {
		node = tempNode
	} else {
		var iterator = node

		for iterator.Next != nil {
			iterator = iterator.Next
		}

		iterator.Next = tempNode
	}

	return node
}

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	var result *utils.ListNode
	var carry int = 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry = carry + l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry = carry + l2.Val
			l2 = l2.Next
		}

		result = add(carry%10, result)
		carry = carry / 10
	}

	if carry > 0 {
		result = add(carry%10, result)
	}

	return result
}

func main() {
	var l1 *utils.ListNode
	var l2 *utils.ListNode

	l1 = utils.MakeSLL([]int{5, 6, 4})
	l2 = utils.MakeSLL([]int{2, 4, 3})

	var result = addTwoNumbers(l1, l2)

	utils.PrintSLL(result)
}
