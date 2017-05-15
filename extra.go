// Write a function to print all integer values in a binary tree in breadth first traversal iteratively. (Hint: Use a Queue.)
// Write a function to print all integer values in a binary tree in pre-order traversal iteratively. (Hint: Use a Stack.)
// How would you implement a Queue using 2 Stacks?
// How would you implement a Stack using 2 Queues?

package main

import (
	"fmt"
)

//Search a sorted array for a value
// func RecursiveBinarySearch(value int, arr []int) bool {
// 	currentIndex := (len(arr) -1)/2 //
// 	currentItem := arr[currentIndex]
// 	subArray := arr

// 	if currentItem == value {
// 		//Return its place
// 		return true
// 	} else if currentItem > value {
// 		subArray = subArray[:currentIndex]
// 		fmt.Println("in here is num is bigger than value")
// 		fmt.Println(subArray)
// 		RecursiveBinarySearch(value, subArray)
// 	} else if currentItem < value {
// 		subArray = subArray[currentIndex:]
// 		fmt.Println("in here is num is smaller than value")
// 		fmt.Println(subArray)
// 		RecursiveBinarySearch(value, subArray)
// 	}

// 	return false
// }

func IterativeBinarySearch(value int, arr []int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (high + low) / 2

		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			low = mid + 1
		} else if arr[mid] > value {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	startArray := []int{1,2,3,4,5,6}
	result := IterativeBinarySearch(5, startArray)
	fmt.Println(result)
}