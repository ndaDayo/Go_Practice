package main

import "fmt"

func main() {
	m := map[string]int{"x": 10, "y": 20}

	fmt.Println(m["x"])

	// 追加する
	m["z"] = 1000

	n, ok := m["z"]
	// zがkeyが存在するかをbool値で返す
	fmt.Println(n, ok)

	// 削除
	delete(m, "z")
}
