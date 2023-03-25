/*
删除排序链表中的重复元素
*/
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func delSortSame(head *ListNode) *ListNode {
	//判断链表或head.Next是否为空 直接返回head 即出现了空表或只包含单个元素的表
	if head == nil || head.Next == nil {
		return head
	}
	//不为空的情况 指定一个指针指向头节点
	currNode := head
	for currNode != nil && currNode.Next != nil {
		if currNode.Val == currNode.Next.Val {
			//如果当前节点的值和下一个节点的值相等 就把下个节点指向下下个节点
			currNode.Next = currNode.Next.Next
		} else {
			//如果不等 就把当前节点指向下一个节点的元素 依次前进
			currNode = currNode.Next
		}
	}
	return head
}
