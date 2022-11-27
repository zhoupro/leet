// Swap Nodes in Pairs
//
// [Medium] [AC:71.2% 546.1K of 767K] [filetype:golang]
//
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
// 示例 1：
//
// 输入：head = [1,2,3,4]
//
// 输出：[2,1,4,3]
//
// 示例 2：
//
// 输入：head = []
//
// 输出：[]
//
// 示例 3：
//
// 输入：head = [1]
//
// 输出：[1]
//
// 提示：
//
// 链表中节点的数目在范围 [0, 100] 内
//
// 0 <= Node.val <= 100
//
// [End of Description]
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main
import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}
	tmp := &ListNode{Val: 5}
	head.Val = 10
	head.Next = tmp
	newHead := swapPairs(head)
	for newHead != nil {
		fmt.Printf("%v\n", newHead.Val)
		newHead = newHead.Next
	}

}

// local end
func swapPairs(head *ListNode) *ListNode {
	listLen := ListLen(head)
	if listLen <= 1 {
		return head
	}

	l1 := head
	l2 := head.Next
	l2Next := l2.Next

	l2.Next = l1
	l1.Next = swapPairs(l2Next)

	return l2
}

func ListLen(head *ListNode) int {
	i := 0
	for head != nil {
		i++
		head = head.Next
	}
	return i

}
