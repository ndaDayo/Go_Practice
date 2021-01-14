package main

import "fmt"

func main() {
	a := []int{10, 20}
	fmt.Println(a, cap(a))

	b := append(a, 30)
	a[0] = 1000
	fmt.Println(b, cap(b))

	fmt.Println(a)

	c := append(b, 40)
	fmt.Println(c, cap(c))

	c[3] = 200
	fmt.Println(c)
}
