package main

import (
	"fmt"
)

func palindrome(text string) bool {
	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Enter word:")
	var text string
	fmt.Scanln(&text)
	fmt.Println(palindrome(text))
}
