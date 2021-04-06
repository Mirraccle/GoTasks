package main

import (
	"fmt"
)

func main() {
	var number, evenSum, oddSum int

	fmt.Print("Enter number: ")
	fmt.Scanln(&number)

	evenSum = 0
	oddSum = 0

	for x := 1; x <= number; x++ {
		if x%2 == 0 {
			evenSum = evenSum + x
		} else {
			oddSum = oddSum + x
		}
	}
	fmt.Println("Sum of even numbers: ", evenSum)
	fmt.Println("Sum of odd numbers: ", oddSum)
}
