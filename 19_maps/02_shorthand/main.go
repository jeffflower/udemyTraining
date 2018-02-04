package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	// adding entry

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)

	// length
	fmt.Println(len(myGreeting))

	// Update

	myGreeting["Harleen"] = "G'day"
	fmt.Println(myGreeting)

	// Delete

	delete(myGreeting, "Harleen")
	fmt.Println(myGreeting)
}
