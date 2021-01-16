package main 

import "fmt"

func main() {
	msg := "Hello, World"
	func ()  {
		fmt.Println(msg)
	}()
}