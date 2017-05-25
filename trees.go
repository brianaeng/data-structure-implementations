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

// Print all values in a Binary Tree using Post-order traversal - left, right, root
func (tree *Tree) Postorder(node *Node) {
	if node == nil {
		return
	}

	tree.Postorder(node.left)

	tree.Postorder(node.right)

	fmt.Println(node.value)
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

//Iterative versions
func (tree *Tree) IterativePreorder(node *Node) {
	
}

func (tree *Tree) IterativePostorder(node *Node) {
	
}

func (tree *Tree) IterativeInorder(node *Node) {
	
}

// Breadth first traversal
func (tree *Tree) BreadthFirst(node *Node) {
	currentNode := tree.root
	finalqQueue := []int

	finalQueue = append(finalQueue, currentNode)

	for len(finalQueue) > 0 {
		if currentNode.left != nil {
			finalQueue = append(finalQueue, currentNode.left)
		}

		if currentNode.right != nil {
			finalQueue = append(finalQueue, currentNode.right)
		}
	}	

// Find the height of a tree (recursion)
func (tree *Tree) FindHeight(node *Node) int {
	if node == nil {
		return 0
	} else {
		leftDepth := FindHeight(node.left)
		rightDepth := FindHeight(node.right)

		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
}

// Find the height of a tree (iterative) -- no idea if this works
func (tree *Tree) FindHeightIterative(node *Node) int {
	queue := []Node
	var currentNode *Node

	queue = append(queue, node)
	height := 0

	for {
		if len(queue) < 1 {
			return height
		}

		for i := 0; i < len(queue); i++ {
			newQueue := []*Node
			currentNode = queue[i]

			if currentNode.left != nil {
				newQueue = append(NewQueue, currentNode.left)
			}

			if currentNode.right != nil {
				newQueue = append(NewQueue, currentNode.right)
			}

		}

		queue = newQueue
	}
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