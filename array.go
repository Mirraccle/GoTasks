package main

import "fmt"

func duplicate(arr []int) int {
	current := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if current[arr[i]] == true {
			return arr[i]
		} else {
			current[arr[i]] = true
		}
	}
	return -1
}

func main() {
	fmt.Println(duplicate([]int{1, 4, 7, 2, 2}))
	fmt.Println(duplicate([]int{1, 4, 7, 2, 3}))
}
