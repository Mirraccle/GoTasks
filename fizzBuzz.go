package main

import (
	"fmt"
)

// func main() {
// 	fmt.Print("Enter number: ")
// 	var n int
// 	fmt.Scanln(&n)

// 	for i := 1; i <= n; i++ {
// 		if i%3 == 0 {
// 			if i%5 == 0 {
// 				fmt.Println(i, ": FizzBuzz")
// 			} else {
// 				fmt.Println(i, ": Fizz")
// 			}
// 		} else if i%5 == 0 {
// 			fmt.Println(i, ": Buzz")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func main() {
// 	fmt.Print("Enter number: ")
// 	var n int
// 	fmt.Scanln(&n)

// 	for i := 1; i <= n; i++ {
// 		if i%15 == 0 {
// 			fmt.Println(i, ": FizzBuzz")
// 			continue
// 		}
// 		if i%5 == 0 {
// 			fmt.Println(i, ": Buzz")
// 			continue
// 		}
// 		if i%3 == 0 {
// 			fmt.Println(i, ": Fizz")
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }

func main() {
	fmt.Print("Enter number: ")
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		switch true {
		case i%15 == 0:
			fmt.Println(i, ": FizzBuzz")
		case i%5 == 0:
			fmt.Println(i, ": Buzz")
		case i%3 == 0:
			fmt.Println(i, ": Fizz")
		default:
			fmt.Println(i)
		}
	}
}
