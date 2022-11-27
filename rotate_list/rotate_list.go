// Rotate List
//
// [Medium] [AC:41.6% 288.9K of 694.7K] [filetype:golang]
//
// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], k = 2
//
// 输出：[4,5,1,2,3]
//
// 示例 2：
//
// 输入：head = [0,1,2], k = 4
//
// 输出：[2,0,1]
//
// 提示：
//
// 链表中节点的数目在范围 [0, 500] 内
//
// -100 <= Node.val <= 100
//
// 0 <= k <= 2 * 109
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
	head := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	head.Next = second

	head = rotateRight(head, 1)
	for head != nil {
		fmt.Printf("%v\n", head.Val)
		head = head.Next
	}

}

// local end

func rotateRight(head *ListNode, k int) *ListNode {
	ll := listLen(head)
    if ll <=1{
        return head
    }
	modL := k % ll
    if modL == 0 {
        return head
    }

    // last point to head
    oriHead := head
    lastNode := &ListNode{}

    for head != nil{
        if head.Next == nil{
            lastNode = head
            lastNode.Next = oriHead
            break
        }
        head = head.Next
    }


    breakHead := oriHead

    for i:=0; oriHead != nil;i++{
        if i == (ll-modL){
            breakHead = oriHead
            break
        }
        oriHead = oriHead.Next
    }

    retHead := breakHead

    for breakHead!=nil{
        if breakHead.Next == retHead{
            breakHead.Next = nil
            break
        }
        breakHead = breakHead.Next
    }

	return retHead

}

func listLen(head *ListNode) int {
	i := 0
	for head != nil {
		i++
		head = head.Next
	}
	return i
}
