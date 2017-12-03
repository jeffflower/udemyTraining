package main

import "fmt"

func foo(n ...int) {
	var sum int
	for _, i := range n {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
