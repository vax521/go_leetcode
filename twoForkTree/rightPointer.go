package main

/*给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
int val;
Node *left;
Node *right;
Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。
解题思路：

1，树问题均可递归

2，先填充一层，然后递归处理左右子树

3，由于是完美二叉树，处理简单：

左子树的next是右子树，右子树的next是next的左子树

*/

type TreeNodeWithNext struct {
	Val   int
	Left  *TreeNodeWithNext
	Right *TreeNodeWithNext
	Next  *TreeNodeWithNext
}

func connect(root *TreeNodeWithNext) *TreeNodeWithNext {
	if root == nil || root.Left == nil {
		return root
	}
	root.Left.Next = root.Right //左子树的next是右子树
	if root.Next != nil {
		root.Right.Next = root.Next.Left //右子树的next是右子树
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
