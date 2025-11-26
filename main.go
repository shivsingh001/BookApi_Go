package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4, 5}

	fmt.Println("length of the array is ", len(num))
	fmt.Println("capacity of the array is ", cap(num))

	num = append(num, 6)

	fmt.Println("length of the array is ", len(num))
	fmt.Println("capacity of the array is ", cap(num))

}

// boring breaks
// do one thing at a time like only eat while eating
// don't use phone for the first 3 hours after waking up
// dopamine stackinng - use it for adding hard work task with fun task
