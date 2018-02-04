package main

import "fmt"

func main() {
	// shorthand creation, creates underlying array with nothing in it
	slice1 := []string{}
	fmt.Println(slice1)
	// var creates zero value slice, no underlying array
	var slice2 []string
	fmt.Println(slice2)
	// idiomatic 2nd parameter is the length, 3rd parameter capacity
	myslice := make([]int, 1)
	// multi dimensional slice
	myslices := make([][]int, 2)
	fmt.Println(myslice[0])
	myslice[0] = 7
	fmt.Println(myslice[0])
	fmt.Println("Capacity:", cap(myslice), "Length:", len(myslice))
	// adding to a slice if length is full
	myslice = append(myslice, 8)
	fmt.Println(myslice[1])
	// Capacity doubles when appending past the capacity amount
	fmt.Println("Capacity:", cap(myslice), "Length:", len(myslice))
	myslice = append(myslice, 10)
	// Capcity doubles again
	fmt.Println("Capacity:", cap(myslice), "Length:", len(myslice))
	// Need to append even though capacity is 4
	myslice = append(myslice, 14)
	fmt.Println("Capacity:", cap(myslice), "Length:", len(myslice))
	// Add 1 to number at index 0
	myslice[0]++
	fmt.Println(myslice[0])
	myslices[0] = myslice
	fmt.Println(myslices)
	// Slice of slice up to index 2
	fmt.Println(myslice[:2])
	// Slice of slice from index 2
	fmt.Println(myslice[2:])
	// Delete item at index 2 from slice
	myslice = append(myslice[:2], myslice[3:]...)
	fmt.Println(myslice)
}
