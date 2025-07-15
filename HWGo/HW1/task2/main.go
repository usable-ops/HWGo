package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

func main() {
	nums1 := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(nums1))

	nums2 := []int{1, 1, 1, 1}
	fmt.Println(numIdenticalPairs(nums2))

	nums3 := []int{1, 2, 3, 1}
	fmt.Println(numIdenticalPairs(nums3))
}
