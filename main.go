package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	num := 1
	for num < 100 {
		num += num
	}
	fmt.Println(num)
}
