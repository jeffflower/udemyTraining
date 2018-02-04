package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		1: "Good morning!",
		2: "Bonjour!",
		3: "Guten Tag",
		4: "Yo!",
	}

	for key, value := range myGreeting {
		fmt.Println(key, " - ", value)
	}
}
