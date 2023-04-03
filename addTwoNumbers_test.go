package main

import (
	"testing"
)

func TestZeroValue(t *testing.T) {
	value1 := &ListNode{0, nil}
	value2 := &ListNode{0, nil}
	expected := &ListNode{0, nil}
	result := AddTwoNumbers(value1, value2)
	if result.Val != expected.Val && result.Next != expected.Next {
		t.Error("unexpected sum result")
	}
}

func TestLeetCodeFirstCase(t *testing.T) {
	value1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	value2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	expected := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
	result := AddTwoNumbers(value1, value2)
	if result.Val != expected.Val {
		t.Error("unexpected result")
	}
}

func TestLeetCodeSecondCase(t *testing.T) {
	value1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	value2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	expected := &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}
	result := AddTwoNumbers(value1, value2)
	for {
		if expected.Val != result.Val {
			t.Error("unexpected result")
			break
		}
		if expected.Next == nil && result.Next == nil {
			break
		}
		expected = expected.Next
		result = result.Next
	}
}
