# Lowest Common Ancestor in a Binary Tree

> Striver SDE Sheet 120 · Day 18 Binary Tree II — recurse both sides; the first node where the targets surface from different subtrees is the answer. Go, stdlib only, tested.

## Problem

Given two nodes of a binary tree (no BST ordering), find their lowest common ancestor — the deepest node having both as descendants, where a node counts as its own descendant (so an ancestor-descendant pair's LCA is the ancestor).

## Approach

Ask each subtree: "did you find p or q?" The recursion returns a found target (or nil), and three outcomes decide everything. Both sides non-nil: the targets split here — this node is the LCA, and it propagates unchanged to the top. One side non-nil: pass it up. The subtle elegance is the base case `root == p || root == q` returning immediately *without looking deeper* — if the other target is below, this node is the LCA anyway (ancestor-descendant case); if not, the other side of some ancestor will produce the second non-nil. One pass, no parent pointers, no paths stored.

**Complexity:** O(n) time, O(h) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
