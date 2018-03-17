package main

import "fmt"

func main() {
	fmt.Println(average(10, 23, 56, 12, 56))
}
func average(nums ...float64) float64 {
	var total float64
	for _, num := range nums {
		// fmt.Println(i, num)
		total += num
	}
	return total / float64(len(nums))
}
