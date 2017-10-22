package main

import "fmt"

func wrapper() func() int {
	x := 0
	fmt.Println(x)
	return func() int {
		x++
		return x
	}
}

func main() {
	fmt.Println("hello1")
	increment := wrapper()
	fmt.Println("hello2")
	fmt.Println(increment())
	fmt.Println(increment())
}
