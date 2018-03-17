package main

import "fmt"

func half(nums ...int) int {
	max := 0
	fmt.Printf("%T", nums)
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
func main() {
	fmt.Println(half(1, 34, 21, 32, 13))
}
