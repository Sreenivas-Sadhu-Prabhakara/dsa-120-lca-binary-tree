package main

import "testing"

func TestLCA(t *testing.T) {
	n4 := tn(4, nil, nil)
	n5 := tn(5, nil, nil)
	n2 := tn(2, n4, n5)
	n6 := tn(6, nil, nil)
	n3 := tn(3, n6, nil)
	root := tn(1, n2, n3)

	cases := []struct {
		p, q, want *TreeNode
	}{
		{n4, n5, n2},   // siblings
		{n4, n6, root}, // across the root
		{n2, n4, n2},   // ancestor-descendant: the ancestor is the LCA
		{n4, n4, n4},   // same node
		{root, n5, root},
	}
	for _, c := range cases {
		if got := LCA(root, c.p, c.q); got != c.want {
			t.Errorf("LCA(%d,%d) = %v, want node %d", c.p.Val, c.q.Val, got, c.want.Val)
		}
	}
}
