package main

import "fmt"

func main() {
	data := []int{23, 2, 2, 7, 2, 5, 9, 1, 4, 317, 12, 5, 8, 12315, 61, 8, 9, 13347, 8, 212, 3412}
	fmt.Printf("Before %v\n", data)
	bubbleSort(data)
	fmt.Printf("After: %v", data)
}

func bubbleSort(in []int) {
	// Start the loop in reverse order, so the loop will start with length
	for i := len(in); i > 0; i-- {
		for j := 1; j < i; j++ {
			if in[j-1] > in[j] {
				in[j], in[j-1] = in[j-1], in[j]
			}
		}
	}
}
