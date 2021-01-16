package main

import "fmt"

func main() {
	fs := make([]func(), 3)
	for i := range fs {
		i := i
		fs[i] = func() { fmt.Println(i) }
	}
	for _, f := range fs {
		f()
	}
}
