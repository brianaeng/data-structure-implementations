// Implement Front for a Queue using only Dequeue, Enqueue(e) & Empty. (Note: This is an example of extending functionality without knowing the implementation details of how the Queue is implemented i.e. not knowing whether the Queue is implemented using an array or a linked list) Assume the data contained is int. Feel free to use additional data structures.
// Empty() which returns true is the data structure is empty, false otherwise
// Min() which returns the minimum integer data value in the data structure
// Max() which returns the maximum integer data value in the data structure

package main

type Queue struct {
	head int
	tail int
	x []int
	size int
}

// Size() which returns the number of items in the data structure
func (queue *Queue) Size() int {
	return queue.size
}

// Front() which returns the item that would be dequeued next (for Queue) 
func (queue *Queue) Front() int {
	return queue.x[head]
}

func (queue *Queue) Dequeue() *Queue {
	queue.x[head] = nil

	head++
	size--

	return queue
} 

func (queue *Queue) Enqueue(value int) *Queue {
	//Remember to change size of array if we add something that pushes size over
}

func (queue *Queue) Empty() bool {

}

