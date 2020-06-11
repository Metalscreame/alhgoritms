package main

import "fmt"

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	return s
}

func main() {

	fmt.Println(string(reverse([]byte("Dogo"))))
}
