package main

import "fmt"

func main2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, "-", j)
		}
	}
}
