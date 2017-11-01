package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("What is up Daniel")
	case "Medhi":
		fmt.Println("What is up Medhi")
	case "Jenny":
		fmt.Println("What is up Jenny")
	default:
		fmt.Println("You have no friends")

	}
}
