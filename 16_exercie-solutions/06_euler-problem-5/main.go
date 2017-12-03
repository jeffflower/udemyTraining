package main

import "fmt"

func main() {
	for i := 20; i < 100000000000; i++ {
		if i%20 == 0 && i%19 == 0 && i%18 == 0 && i%17 == 0 && i%16 == 0 && i%14 == 0 && i%13 == 0 && i%11 == 0 {
			fmt.Println(i)
			break
		}
	}
}

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
