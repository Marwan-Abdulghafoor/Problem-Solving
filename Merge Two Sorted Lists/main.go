package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list2 == nil {
		return list1
	}
	if list1 == nil {
		return list2
	}
	var cur1 *ListNode
	var cur2 *ListNode
	var head *ListNode
	var temp *ListNode
	if list1.Val < list2.Val || list1.Val == list2.Val {
		head = list1
		temp = list1
		cur1 = list1.Next
		cur2 = list2
	} else {
		head = list2
		temp = list2
		cur1 = list1
		cur2 = list2.Next
	}
	for cur1 != nil || cur2 != nil {
		if cur2 == nil {
			temp.Next = cur1
			temp = cur1
			cur1 = cur1.Next
		} else if cur1 == nil {
			temp.Next = cur2
			temp = cur2
			cur2 = cur2.Next
		} else if cur1.Val < cur2.Val || cur1.Val == cur2.Val {
			temp.Next = cur1
			temp = cur1
			cur1 = cur1.Next
		} else {
			temp.Next = cur2
			temp = cur2
			cur2 = cur2.Next
		}
	}
	return head
}
