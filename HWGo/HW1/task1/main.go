package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 5}
	sLength, n := len(slice), len(slice)
	newSize := 2 * sLength
	newSlice := make([]int, newSize)
	for i := 0; i < sLength; i++ {
		newSlice[i], newSlice[i+n] = slice[i], slice[i]
	}
	fmt.Println(newSlice)
}
