package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		1: "Good morning!",
		2: "Bonjour!",
		3: "Guten Tag",
		4: "Yo!",
	}

	fmt.Println(myGreeting)

	if val, exists := myGreeting[2]; exists {
		delete(myGreeting, 2)
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}

	fmt.Println(myGreeting)
}
