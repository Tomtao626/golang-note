package main

type ListNode struct {
	Val  string
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	var listNode *ListNode
	listNode.Val = "1,2,3,4,5,6"
	reverseList(listNode)
}
