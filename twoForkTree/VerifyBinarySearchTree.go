package main

import "fmt"

/**
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：
节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:
输入:
    2
   / \
  1   3
输出: true
示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var last = ^(int(^uint(0) >> 1)) //表示一个极小值

/**
思路一：中序遍历
*/
func isvalidbst1(root *TreeNode) bool {
	if root != nil {
		fmt.Println(last)
		fmt.Println(root.Val)
		fmt.Println()
		if !isvalidbst1(root.Left) {
			return false
		}
		if last >= root.Val {
			return false
		}
		last = root.Val
		if !isvalidbst1(root.Right) {
			return false
		}
	}
	return true
}

func maxBST(root *TreeNode) int {
	if root.Right != nil {
		return maxBST(root.Right)
	}
	return root.Val
}

func minBST(root *TreeNode) int {
	if root.Left != nil {
		return minBST(root.Left)
	}
	return root.Val
}

/**
  递归：根节点>大于左节点最大值，小于右节点最小值
*/
func isvalidbst2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Right != nil {
		maxLeft := maxBST(root.Left)
		minRight := minBST(root.Right)
		return isvalidbst2(root.Left) && isvalidbst2(root.Right) && maxLeft < root.Val && minRight > root.Val
	}
	if root.Left != nil {
		maxLeft := maxBST(root.Left)
		return isvalidbst2(root.Left) && maxLeft < root.Val
	}
	if root.Right != nil {
		minRight := minBST(root.Right)
		return isvalidbst2(root.Right) && minRight > root.Val
	}
	return true
}
func main() {
	node1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	//fmt.Println(isValidBST_1(node1))
	fmt.Println(isvalidbst2(node1))
}
