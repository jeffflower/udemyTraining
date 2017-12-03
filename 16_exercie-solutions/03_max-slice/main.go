package main

import "fmt"

func maxslice(sf ...float64) (max float64) {
	for _, v := range sf {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	n := maxslice(43, 65, 78, 12, 235, 321)
	fmt.Println(n)
}
