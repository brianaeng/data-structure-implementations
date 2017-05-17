// Find the height of Binary Tree. Implement the recursive solution.
// Insert a given value in a Binary Search Tree. Implement recursive solution.
// Find a given value in a Binary Search Tree. Implement recursive solution. 

//also do all the problems iteratively

//insert the nodes and then print them out in the right order

//given a string, from string & to string (of same length) - first char of from string should be replaced by from
//find & replace given to and from

//operating system concepts - dinosaur book
//dining philosopher

//Thursday after class --trees depth first traversal (preorder, postorder, etc.)
//convert stacks/queue to linked list

//cs stanford

//string manipulation using bit arrays

package main

import (
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	value int
	left *Node
	right *Node
}

// Print all values in a Binary Tree using Pre-order traversal - root, left, right
func (tree *Tree) Preorder(node *Node) {
	if node == nil {
		return 
	}

	fmt.Println(node.value)

	tree.Preorder(node.left)

	tree.Preorder(node.right)
}

func (tree *Tree) IterativePreorder(node *Node) {
	
}

// Print all values in a Binary Tree using Post-order traversal - left, right, root
func (tree *Tree) Postorder(node *Node) {
	if node == nil {
		return
	}

	tree.Postorder(node.left)

	tree.Postorder(node.right)

	fmt.Println(node.value)
}

func (tree *Tree) IterativePostorder(node *Node) {
	
}

// Print all values in a Binary Tree using In-order traversal - left, root, right
func (tree *Tree) Inorder(node *Node) {
	if node == nil {
		return
	}

	tree.Inorder(node.left)

	fmt.Println(node.value)

	tree.Inorder(node.right)
}

func (tree *Tree) IterativeInorder(node *Node) {
	
}

func (tree *Tree) DepthFirst(node *Node) {

}

func (tree *Tree) BreadthFirst(node *Node) {

}

// Delete a given value from a Binary Search Tree using recursive solution.
func (tree *Tree) Delete(value int) bool {
	currentNode := tree.root

	return tree.RemoveNode(currentNode, value)
}

func (tree *Tree) RemoveNode(node *Node, value int) bool {
	if node == nil {
		return node
	}

	if value < node.value {
		node.left = tree.RemoveNode(node.left, value)
	} else if value > node.value {
		node.right = tree.RemoveNode(node.right, value)
	} else {
		if node.left == nil {

		}
	}
	
}