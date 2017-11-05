package main

import "fmt"

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()   // x is created here
	fmt.Println(increment()) // func is called, x is incremented
	fmt.Println(increment()) // func is called, x is incremented
}
