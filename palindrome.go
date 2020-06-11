package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 { // can't be palindrome
		return false
	}

	reversed := 0
	tmp := x

	for tmp > 0 {
		reversed = reversed*10 + tmp%10
		tmp /= 10
	}

	return reversed == x
}

func main() {
	fmt.Println(isPalindrome(121))
}
