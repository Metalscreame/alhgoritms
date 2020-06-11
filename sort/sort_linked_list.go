package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func prepeareLinkedList() *ListNode {
	head := ListNode{Val: 1}
	first := ListNode{Val: 6}
	second := ListNode{Val: 2}
	third := ListNode{Val: 4}
	forth := ListNode{Val: 9}

	head.Next = &first
	first.Next = &second
	second.Next = &third
	third.Next = &forth
	return &head
}

func sortList(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return head
	}

}
