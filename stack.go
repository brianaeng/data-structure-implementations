// Implement Top for a Stack using only Push(e) and Pop. (Note: This is an example of extending functionality without knowing the implementation details of how the Stack is implemented i.e. not knowing whether the Stack is implemented using an array or a linked list). Assume the data contained is int.
// Top() which returns the item that would be popped next  (for Stack)

package main 

type Stack struct {
	x []int
	size int
}

func (stack *Stack) Top() int {
	return stack.x[0]

	//way of doing this with only pop & push:
}

func (stack *Stack) Empty() bool {
	if stack.size == 0 {
		return true
	}

	return false
}

// Write a function to reverse a string (the funciton takes a string as input parameter) using a Stack.
func (stack *Stack) Reverse(str string) string {

}

