package main

import "fmt"

func main() {

	for i, v := range []int{1, 2, 3} {
		fmt.Println(i)
		fmt.Println(v)
	}
}
