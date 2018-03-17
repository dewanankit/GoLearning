package main

import "fmt"

func greeter() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
func main() {
	greetVar := greeter()
	fmt.Println(greetVar())
	fmt.Println(greetVar())
}
