package main

import "fmt"

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	sum := AddTwoNumbers(l1, l2)
	for {
		fmt.Print(sum.Val)
		if sum.Next == nil {
			break
		}
		sum = sum.Next
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	overhead := 0

	for l1 != nil || l2 != nil {
		firstDigit := 0
		secondDigit := 0
		if l1 != nil {
			firstDigit = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			secondDigit = l2.Val
			l2 = l2.Next
		}
		sum := firstDigit + secondDigit + overhead
		overhead = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	if overhead > 0 {
		current.Next = &ListNode{Val: overhead}
	}

	return result.Next
}
