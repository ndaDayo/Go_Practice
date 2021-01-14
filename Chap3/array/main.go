package main

import "fmt"

// 固定長配列なので、いちいち長さを書かないとダメ 
var array = [5]int{1, 2, 3, 4, 5}

func main() {
	fmt.Println(array)
	fmt.Println(array[1])
	
	// ... を書くと、推論してくれる。便利
	array2 := [...]int{10,20,30}
	fmt.Println(array2)
}
