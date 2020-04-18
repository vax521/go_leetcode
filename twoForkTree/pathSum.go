package main

import "fmt"

/*
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

解题思路：

1，对于二叉树类型题目一般都是递归解

2，递归有两种：自根向下和自叶子向上

3，对于相等类型题目和最大值题目解题思路相反，本题采用自根向下
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		return root.Val == sum
	}
	return findPathSum(root.Left, sum-root.Val) || findPathSum(root.Right, sum-root.Val)
}

func findSumWayPath(root *TreeNode, sum int) [][]int {
	var r [][]int
	if root == nil {
		return r
	}
	var temp []int
	return walk(root, sum, temp)
}

func walk(root *TreeNode, sum int, temp1 []int) [][]int {
	var r [][]int
	if root == nil {
		return r
	}
	temp := make([]int, len(temp1))
	copy(temp, temp1)
	temp = append(temp, root.Val)
	fmt.Println(root, temp, sum)
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			r = append(r, temp)
			// fmt.Println(root,temp,sum,r)
			return r
		}
		return r
	}

	tl := walk(root.Left, sum-root.Val, temp)
	tr := walk(root.Right, sum-root.Val, temp)

	if len(tl) > 0 {
		r = append(r, tl...)
	}
	if len(tr) > 0 {
		r = append(r, tr...)
	}
	return r
}

func main() {

	node9 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	node8 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	node7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	node6 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: node9,
	}

	node5 := &TreeNode{
		Val:   13,
		Left:  nil,
		Right: nil,
	}

	node4 := &TreeNode{
		Val:   11,
		Left:  node7,
		Right: node8,
	}

	node3 := &TreeNode{
		Val:   8,
		Left:  node5,
		Right: node6,
	}

	node2 := &TreeNode{
		Val:   4,
		Left:  node4,
		Right: node3,
	}

	node1 := &TreeNode{
		Val:   5,
		Left:  node2,
		Right: node3,
	}

	//fmt.Println(findPathSum(node1,22))
	fmt.Println(findSumWayPath(node1, 22))
}
