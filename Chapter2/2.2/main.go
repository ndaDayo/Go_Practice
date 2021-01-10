package main

import (
	"fmt"
)

// インターフェース
type Talker interface {
	Talk()
}

// 構造体
type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("my name is %s\n", g.name)
}

func main() {
	var talker Talker

	talker = &Greeter{"wozozo"}
	talker.Talk()
}
