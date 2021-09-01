package main

// "Given a binary tree, find the maximum path sum.
// The path may start and end at any node in the tree."

type GraphNode struct {
	value int

	left *GraphNode
	right *GraphNode
}

