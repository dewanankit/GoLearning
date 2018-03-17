package main

import "fmt"

func tester(nums []int, callback func(a int)) {
	for _, n := range nums {
		callback(n)
	}
}
func main() {
	tester([]int{1, 2, 3, 4, 5}, func(a int) {
		fmt.Println(a)
	})
}
