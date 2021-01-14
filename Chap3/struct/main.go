package main

import "fmt"

func main() {
	p := struct {
		name string
		age  int
	}{name: "Nda", age: 33}

	fmt.Println(p.age)
}
