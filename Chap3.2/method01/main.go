package main

import "fmt"

func main() {
	fs := make([]func() string, 2)
	fs[0] = func() string { return "nda" }
	fs[1] = func() string { return "dayo" }
	for _, f := range fs {
		fmt.Println(f())
	}
}
