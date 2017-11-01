package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("The sum of natural numbers between 1 and 1000 that are multiples of 3 or 5 is:", sum)
}
