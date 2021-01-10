package main

import (
	"fmt"
)

var pow = []int{2, 3, 4, 5, 13, 43, 8}

func main() {
	for i, v := range pow {
		fmt.Println(i)
		fmt.Println(v)
	}
}
