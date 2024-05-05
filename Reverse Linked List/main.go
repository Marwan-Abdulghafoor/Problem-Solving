package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head.Next
	head.Next = nil
	for cur != nil {
		temp := cur.Next
		cur.Next = head
		head = cur
		cur = temp
	}
	return head
}
