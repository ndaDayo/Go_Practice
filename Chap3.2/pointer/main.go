package main

import (
	"fmt"
)

func f(xp *int) {
	// *xp はポインタの指す先 ポインタは変数が格納される場所
	// xp のポインタは0xc0000ba018　値は0xc000016050
	*xp = 100
}

func main() {
	var x int
	// ＆xはポインタを参照 0xc000016050
	f(&x)
	fmt.Println(x)
}
