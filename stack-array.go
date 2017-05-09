// Implement Top for a Stack using only Push(e) and Pop. (Note: This is an example of extending functionality without knowing the implementation details of how the Stack is implemented i.e. not knowing whether the Stack is implemented using an array or a linked list). Assume the data contained is int.
// Top() which returns the item that would be popped next  (for Stack)

package main 

import (
	"fmt"
)

// type Stack struct {
// 	x []int
// 	size int
// }

type Stack struct {
	x []string
	size int
}

func (stack *Stack) Top() string {
	// return stack.x[0]

	//way of doing this with only pop & push:
	storedValue := stack.Pop()
	stack.Push(storedValue)
	return storedValue
}

func (stack *Stack) Push(value string) *Stack {
	stack.x = append(stack.x, value)
	stack.size++
	return stack
}

func (stack *Stack) Pop() string {
	value := stack.x[stack.size -1]
	stack.x[stack.size -1] = ""
	stack.size--

	return value
}

func (stack *Stack) Empty() bool {
	if stack.size == 0 {
		return true
	}

	return false
}

// Write a function to reverse a string (the funciton takes a string as input parameter) using a Stack.
func (stack *Stack) Reverse(str string) string {
	newString := ""

	for _, char := range str {
		fmt.Println(string(char))
		stack.Push(string(char))
	}

	//Need to hold stack size since we mutate it using Pop()
	stackSize := stack.size

	for i := 0; i < (stackSize); i++ {
		newString += stack.Pop()
	}

	return newString
}

func main() {
	newStack := Stack{x: []string{}, size:0}

	finalString := newStack.Reverse("hello")
	fmt.Println(finalString)
}

