// Find the nth node from the end of a singly linked list, assuming index starting at 0. Return a pointer or reference to it.
// Insert a node into a singly linked list sorted in ascending order of data
// Remove the first 5 nodes from a singly linked list
// Empty a singly linked list
// Empty an array of size 20. Empty = uninitialize all cells to -1
// Insert a node into a sorted singly linked list with value 999.
// Add an element with value 999 to a sorted array. The array is of size 20. All unassigned entries have value INT_MAX, which is specially reserved for this purpose. All unassigned entries are at (or shifted to be at) the end of the array.
// Delete nodes in a singly linked list with specified data value e.g. int value of 5
// Remove entries from an array with specified data value e.g. int value of 5 The array is of size 20 and all unassigned entries have value INT_MAX, which is specially reserved for this purpose. All unassigned entries are at (or shifted to be at) the end of the array.
// Print the largest integer data value in a singly linked list
// Print the largest integer data value in a native int array
// Print all integer data values in a native int array
// Append a node to the beginning of a singly linked list
// Check if the singly linked list contains a given integer value.
// Check if a native integer array contains a given integer value.
// Reverse a native integer array without using helper methods available in higher level programming languages 
// Note: For all singly linked list questions, assume each node in the linked list to have integer data and link to the next node
// Implement a function to add a Node to the beginning of a doubly linked list.
// Implement a function to remove a Node from the beginning of a doubly linked list.
// Implement a function to add a Node at the second position in a doubly linked list.
// Implement a function to remove the second Node from a doubly linked list.

package main 

import (
	"fmt"
)

type LinkedList struct {
  head *Node
  size int
}

type Node struct {
  value int
  next *Node
}

//Find the length of a LinkedList
func (list *LinkedList) FindLength() int {
  return list.size
}

//Add a new Node to the end of a LinkedList
func (list *LinkedList) InsertNode(value int) *LinkedList {
  currentNode := list.head

  for currentNode.next != nil {
    currentNode = currentNode.next
  }

  newNode := Node{value: value, next:nil}

  currentNode.next = &newNode
  list.size++

  list.PrintValues()
  return list
}

//Remove a Node
func (list *LinkedList) RemoveNode(value int) *LinkedList {
  currentNode := list.head
  var previousNode *Node

  for i := 0; i < list.FindLength(); i++ {
    if currentNode.value == value {
      fmt.Println("Broke")
      break
    }
    previousNode = currentNode
    currentNode = currentNode.next
  }

  if currentNode.next == nil {
    fmt.Println("No node with passed value")
    return list
  }

  previousNode.next = currentNode.next
  list.size -= 1

  list.PrintValues()

  return list
}

// Print all integer data values in a singly linked list
func (list *LinkedList) PrintValues() *LinkedList {
  currentNode := list.head

  for currentNode != nil {
    fmt.Println(currentNode.value)
    currentNode = currentNode.next
  }

  return list
}

// Reverse a singly linked list. (Test: if reference to the first node is recorded before the function call, it should be pointing at the last node in the singly linked list after the function call is complete.)
func (list *LinkedList) Reverse() *LinkedList {
  currentNode := list.head
  var nextNode *Node
  var previousNode *Node

  for currentNode != nil {
    nextNode = currentNode.next

    currentNode.next = previousNode 
    previousNode = currentNode
    currentNode = nextNode
  }

  list.head = previousNode

  return list
}

// Check if a singly linked list has a cycle. Return true if it has a cycle, false if it doesn't.
func (list *LinkedList) CycleCheck() bool {
  slow := list.head
  fast := list.head

  for slow != nil && fast != nil && fast.next != nil {
    slow = slow.next
    fast = fast.next.next

    if fast == slow {
      fmt.Println(true)
      return true
    }
  }
  fmt.Println(false)
  return false
}

// Find the middle node in a singly linked list. Return a pointer or reference to it.
func (list *LinkedList) FindMiddle() *Node {
  length := list.FindLength()
  desiredNode := length/2
  currentNode := list.head

  for i := 0; i < desiredNode; i++ {
    currentNode = currentNode.next
  }

  fmt.Println(currentNode.value)
  return currentNode
}

func main() {
  fifthNode := Node{value: 5, next: nil}
  fourthNode := Node{value:4, next: &fifthNode}
  thirdNode := Node{value:3, next: &fourthNode}
  secondNode := Node{value:2, next:&thirdNode}
  firstNode := Node{value:1, next: &secondNode}
  myList := LinkedList{head:&firstNode, size:5}

  myList.FindLength()

  // Testing FindMiddle()
  // myList.FindMiddle()

  // Testing Reverse()
  // myList.Reverse()

  // Testing CycleCheck()
  // fourthNode.next = &secondNode
  // myList.CycleCheck()

  // myList.InsertNode(6)
  // myList.RemoveNode(4)
}