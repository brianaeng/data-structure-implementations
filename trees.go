// Delete a given value from a Binary Search Tree using recursive solution.
// Print all values in a Binary Tree using Pre-order, Post-order and In-order traversals. Implement each solution recursively.
// Find the height of Binary Tree. Implement the recursive solution.
// Insert a given value in a Binary Search Tree. Implement recursive solution.
// Find a given value in a Binary Search Tree. Implement recursive solution. 

package main

type Tree struct {
	root *Node
}

type Node struct {
	value int
	left *Node
	right *Node
}

func (tree *Tree) Delete(value int) *Tree {
	
}