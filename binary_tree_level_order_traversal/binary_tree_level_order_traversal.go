// Binary Tree Level Order Traversal
//
// [Medium] [AC:65.3% 727.9K of 1.1M] [filetype:golang]
//
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。
// （即逐层地，从左到右访问所有节点）。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
//
// 输出：[[3],[9,20],[15,7]]
//
// 示例 2：
//
// 输入：root = [1]
//
// 输出：[[1]]
//
// 示例 3：
//
// 输入：root = []
//
// 输出：[]
//
// 提示：
//
// 树中节点数目在范围 [0, 2000] 内
//
// -1000 <= Node.val <= 1000
//
// [End of Description]
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 10,
	}
	levelOrder(root)

}

// local end

type Stack struct {
	Items []*TreeNode
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) Push(node *TreeNode) {
	s.Items = append(s.Items, node)
}

func (s *Stack) Pop() *TreeNode {
	node := s.Items[0]
	s.Items = s.Items[1:]
	return node
}

func levelOrder(root *TreeNode) [][]int {

	nodes := make([]int, 0)
	retNodes := make([][]int, 0)

	if root == nil {
		return retNodes
	}

	stack1 := &Stack{}
	stack2 := &Stack{}
	stack1.Push(root)

	for !stack1.IsEmpty() || !stack2.IsEmpty() {
		curNode := stack1.Pop()
		nodes = append(nodes, curNode.Val)
		if curNode.Left != nil {
			stack2.Push(curNode.Left)
		}

		if curNode.Right != nil {
			stack2.Push(curNode.Right)
		}
		if stack1.IsEmpty() {
			stack1 = stack2
			stack2 = &Stack{}
			retNodes = append(retNodes, nodes)
			nodes = make([]int, 0)
		}

	}

	return retNodes
}
