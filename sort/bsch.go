package main

import "fmt"

func bsrch(nums []int, target int) int {
	// lets firts check if list is not empty
	if len(nums) == 0 {
		return -1
	}
	// as this aray is sorted we can find value in the miidle and undestand in which part of array
	//  we can find target value. we wil be doing this dividing by two until we find targt value
	// or until we chek all parts of array
	low, high := 0, len(nums)-1 // lets define two indexes of a slice
	mid := 0 //and define mid variable

	for low <= high { // we will be iterating until indexes are in aproptiate states.
	// so until low pos is less or equals to high
	// so first lets find mid value
		mid = (high + low) >> 1 // or (high+low)/2. this is right shift operator
		if nums[mid] == target { // and then lets check if the the value by middle index is the target
			return mid
		}
		// if not lets check whether mid val is less then target value
		// if yes - then we cut first half of an array
		if nums[mid] < target {
			low = mid + 1 // and make low index as mid +1. as mid ind is already checked
			continue
		}
		// if no this means that target is in low part and we need to cut hight part
		high = mid - 1
	}

	// if target is still not found - return that it is not found
	return -1
}

func main() {
	nums := []int{-1, 0, 2, 3, 5, 9, 12}
	target := 12

	index := bsrch(nums, target)
	fmt.Println(index)
}
