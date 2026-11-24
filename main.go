package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// LCA returns the lowest common ancestor of nodes p and q (compared by
// pointer identity; both assumed present). A node is the answer exactly
// when the two targets surface from different sides beneath it.
func LCA(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := LCA(root.Left, p, q)
	r := LCA(root.Right, p, q)
	switch {
	case l != nil && r != nil:
		return root // one target on each side: this is the split point
	case l != nil:
		return l
	default:
		return r
	}
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	n4 := tn(4, nil, nil)
	n5 := tn(5, nil, nil)
	n2 := tn(2, n4, n5)
	n3 := tn(3, nil, nil)
	root := tn(1, n2, n3)
	fmt.Println(LCA(root, n4, n5).Val, LCA(root, n4, n3).Val)
}
