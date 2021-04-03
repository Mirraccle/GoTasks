package main

import (
	"fmt"
)

func main() {
	fmt.Print("Введите номер элемента ряда чисел Фибоначи: ")
	var n int
	fmt.Scanln(&n)

	var fib1 = 0 // Первое число ряда Фибоначи
	var fib2 = 1 // Второе число ряда Фибоначи
	var c int    // Запасная переменная для подмены значений
	fmt.Println(fib1)
	fmt.Println(fib2)

	for i := 3; i <= n; i++ {
		c = fib2
		fib2 = fib1 + fib2
		fib1 = c
		fmt.Println(fib2)
	}
}

// Recursion

// func Fibonacci(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return Fibonacci(n-1) + Fibonacci(n-2)
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Print(Fibonacci(i), "\n")
// 	}
// }
