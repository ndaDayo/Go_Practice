package main

import (
	"fmt"
)

func main() {
	primes := [4]int{2, 3, 4, 5}

	var s []int = primes[0:2]
	fmt.Println(s)
}
