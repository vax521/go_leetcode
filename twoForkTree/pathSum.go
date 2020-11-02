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

type Node struct {
	Val  *TreeNode
	pre  *Node
	next *Node
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

/*
 * 二叉树前序遍历   根-> 左-> 右
 * @param node    二叉树节点
 */

func preOrderTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	preOrderTraveral(node.Left)
	preOrderTraveral(node.Right)
}

/*
 * 二叉树中序遍历    左-> 根->右
 * @param node    二叉树节点
 */

func inOrderTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	preOrderTraveral(node.Left)
	fmt.Printf("%d ", node.Val)
	preOrderTraveral(node.Right)
}

/*
 * 二叉树后序遍历    左->右-> 根
 * @param node    二叉树节点
 */

func postOrderTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	preOrderTraveral(node.Left)
	preOrderTraveral(node.Right)
	fmt.Printf("%d ", node.Val)
}

//Queue 队列
type Queue struct {
	top    *Node
	rear   *Node
	length int
}

// Create a new queue
func NewQueue() *Queue {
	return &Queue{nil, nil, 0}
}

//入队操作
func (this *Queue) Push(treeNode *TreeNode) {
	n := &Node{treeNode, nil, nil}
	if this.length == 0 {
		this.top = n
		this.rear = this.top
	} else {
		n.pre = this.rear
		this.rear.next = n
		this.rear = n
	}
	this.length++
}

//出队操作
func (this *Queue) Pop() *TreeNode {
	if this.length == 0 {
		return nil
	}
	n := this.top
	if this.top.next == nil {
		this.top = nil
	} else {
		this.top = this.top.next
		this.top.pre.next = nil
		this.top.pre = nil
	}
	this.length--
	return n.Val
}

/*
首先申请一个新的队列，记为queue；
将头结点head压入queue中；
每次从queue中出队，记为node，然后打印node值，如果node左孩子不为空，则将左孩子入队；如果node的右孩子不为空，则将右孩子入队；
重复步骤3，直到queue为空。
*/

func levelOrderTraveral(node *TreeNode) {
	queue := NewQueue()
	queue.Push(node)
	for {
		if queue.length == 0 {
			break
		} else {
			root := queue.Pop()
			fmt.Println(root.Val)
			if root.Left != nil {
				queue.Push(root.Left)
			}
			if root.Right != nil {
				queue.Push(root.Right)
			}
		}

	}
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
	//fmt.Println(findSumWayPath(node1, 22))
	preOrderTraveral(node1)
	fmt.Println()
	inOrderTraveral(node1)
	fmt.Println()
	postOrderTraveral(node1)
	fmt.Println()
	levelOrderTraveral(node1)
}
