package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "World"

	fmt.Println(a)

	primes := [6]int{2,3,4,5,6,7}
	fmt.Println(primes)
}
