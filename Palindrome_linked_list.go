package main

import "fmt"

func isPalindromeLinked(head *ListNode) bool {
	if head == nil || head.Next == nil { // can't be palindrome
		return false
	}

	slow, fast := head, head

	// slow is moved x1, fast x2
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = head
	slow = reverseLinkedList(slow)

	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prevNode *ListNode

	if head == nil {
		return head
	}

	for head != nil {
		nextNode := head.Next
		head.Next = prevNode
		prevNode = head
		head = nextNode
	}
	return prevNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	fmt.Println(isPalindromeLinked(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	}))
}
