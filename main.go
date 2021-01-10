package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.5451, -76.5452},
	"Google":    {34.344, 43.3443},
}

func main() {
	fmt.Println(m)
}
