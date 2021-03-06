package main

// Downside of using an array as underlying data strucutre: resizing array would mean making a new one and moving everything over
// How do I resize arrays? Do I keep doubling the size?

// Try implementing with a linked list too
type Queue struct {
	head int //index of first value
	tail int // index of last value
	x []int
	size int
}

// Size() which returns the number of items in the data structure
func (queue *Queue) Size() int {
	return queue.size
}

// Front() which returns the item that would be dequeued next (for Queue) 
func (queue *Queue) Front() int {
	// return queue.x[head]

	//way of doing this with only queue & dequeue:
	storedValue := queue.Dequeue()
	queue.Enqueue(storedValue)

	return storedValue
}

func (queue *Queue) Dequeue() int {
	value := queue.x[head]
	queue.x[head] = nil

	head++
	size--

	return value
} 

func (queue *Queue) Enqueue(value int) *Queue {
	//Remember to change size of array if we add something that pushes size over
	queue.x[tail + 1] = value
	tail++
	size++

	return queue
}

// Thinking I should have a method that checks if a resize is needed & resizes as needed
func (queue *Queue) ResizeCheck() *Queue {
	return queue
}

// Empty() which returns true is the data structure is empty, false otherwise
func (queue *Queue) Empty() bool {
	if size == 0 {
		return true
	}
	return false

	//Or?
	// for i := 0; i < len(q.x); i++ {
	// 	if x[i] != nil {
	// 		return false
	// 	}
	// return true
	}
}

// Min() which returns the minimum integer data value in the data structure
func (queue *Queue) Min() int {
	if queue.Empty() {
		return 0
	}

	minNum := queue.x[head]
	for i := head; i < (queue.size - 1); i++ {
		if queue[i] < minNum {
			minNum = queue[i]
		}
	}

	return minNum
}

// Max() which returns the maximum integer data value in the data structure
func (queue *Queue) Max() int {
	if queue.Empty() {
		return 0
	}

	maxNum := queue.x[head]
	for i := head; i < (queue.size - 1); i++ {
		if queue[i] > maxNum {
			maxNum = queue[i]
		}
	}

	return maxNum
}

