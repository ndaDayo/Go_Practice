package main

import "fmt"

func main() {
	ns := []int{10, 20, 30}
	// ポインタもコピーされる
	ns2 := ns
	// ns[1]とns2[1]のポインタの値は同じ
	ns[1] = 200
	// 0xc0000c0008
	fmt.Println(&ns[1])
	//0xc0000c0008
	fmt.Println(&ns2[1])

	// なので、↓は同じ結果になる。
	fmt.Println(ns[0], ns[1], ns[2])
	fmt.Println(ns2[0], ns2[1], ns2[2])

}
