package main

import "fmt"

/*
https://mp.weixin.qq.com/s/68gMHsGHxxecObC6m1Ug4w
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queue struct {
	data []*TreeNode
}

func (q *queue) last() *TreeNode {
	queueLength := len(q.data)
	var node *TreeNode
	if queueLength > 0 {
		node = q.data[queueLength-1]
	}
	return node
}

func (q *queue) push(node *TreeNode) {
	q.data = append(q.data, node)
}

func (q *queue) pop() *TreeNode {
	queueLength := len(q.data)
	var node *TreeNode
	if queueLength > 0 {
		node = q.data[0]
		q.data = q.data[1:]
	}
	return node
}

func (q *queue) empty() bool {
	return len(q.data) <= 0
}

func (q *queue) getLeftNodeVal() int {
	if len(q.data) > 0 {
		return q.data[0].Val
	} else {
		return -100000
	}
}

/**
  打印每一层的节点值
*/
func print(que queue) {
	if !que.empty() {
		node := que.pop()
		fmt.Println(node.Val)
	}
}

/**
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/
func rightSideView(root *TreeNode) []int {
	var q queue //第一个队列存上一层结果
	var result []int
	if root != nil {
		q.push(root)
	}
	for !q.empty() {
		lastNode := q.last()
		result = append(result, lastNode.Val)
		var q1 queue //第二个队列存上一层的所有孩子node
		for !q.empty() {
			tempNode := q.pop()
			if tempNode.Left != nil {
				q1.push(tempNode.Left)
			}
			if tempNode.Right != nil {
				q1.push(tempNode.Right)
			}
		}
		q = q1
	}
	return result
}

/**
给定一个二叉树，在树的最后一行找到最左边的值。
示例 1:
输入:

    2
   / \
  1   3
输出:
1

示例 2:

输入:

        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7

输出:
7
*/

func lastLeftVal(root *TreeNode) int {
	var q queue //第一个队列存上一层结果
	var result int
	if root != nil {
		q.push(root)
	}
	for !q.empty() {
		var q1 queue //第二个队列存上一层的所有孩子node
		result = q.getLeftNodeVal()
		for !q.empty() {
			tempNode := q.pop()
			if tempNode.Left != nil {
				q1.push(tempNode.Left)
			}
			if tempNode.Right != nil {
				q1.push(tempNode.Right)
			}
		}
		q = q1
	}
	return result
}

/**
您需要在二叉树的每一行中找到最大的值。

示例：

输入:

          1
         / \
        3   2
       / \   \
      5   3   9

输出: [1, 3, 9]
*/
func getMaxEveryLine(root *TreeNode) []int {
	var q queue //第一个队列存上一层结果
	var result []int
	if root != nil {
		q.push(root)
	}
	for !q.empty() {
		var q1 queue //第二个队列存上一层的所有孩子node
		max := q.getLeftNodeVal()
		for !q.empty() {
			tempNode := q.pop()
			if tempNode.Val > max {
				max = tempNode.Val
			}
			if tempNode.Left != nil {
				q1.push(tempNode.Left)
			}
			if tempNode.Right != nil {
				q1.push(tempNode.Right)
			}
		}
		q = q1
		result = append(result, max)
	}
	return result
}

/**
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明: 叶子节点是指没有子节点的节点。
示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
[
   [5,4,11,2],
   [5,8,4,5]
]
*/

func main() {
	node5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: node4,
	}
	node2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: node5,
	}
	node1 := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}
	//fmt.Println(rightSideView(node1))
	//fmt.Println(lastLeftVal(node1))
	fmt.Println(getMaxEveryLine(node1))
}
