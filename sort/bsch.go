package main

import "fmt"

func bsrch(nums []int, target int) int {
	// lets firts check if list is not empty
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	mid := 0
	for low <= high {
		mid = (high + low) >> 1 // or (high+low)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
			continue
		}

		high = mid - 1
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 2, 3, 5, 9, 12}
	target := 12

	index := bsrch(nums, target)
	fmt.Println(index)
}
