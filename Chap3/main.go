package main

import "fmt"

func main() {
	var sum int
	sum = 1 + 2 + 3
	avg := float64(sum) / 3
	if avg > 0.5 {
		fmt.Println("good")
	}
}
