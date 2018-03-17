package main

import "fmt"

func main() {

	greet("John", "doe")
	fmt.Println(greet("Jane", "hoe"))
}
func greet(name, lastname string) (string, string) {
	return fmt.Sprint(name, lastname), fmt.Sprint(lastname, name)
}
