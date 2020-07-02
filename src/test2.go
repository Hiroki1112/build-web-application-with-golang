package main

import "fmt"

func main() {
	s := "hello"
	c := []byte(s) // 文字列 s を []byte 型にキャスト
	c[0] = 'c'
	s2 := string(c) // もう一度 string 型にキャストし直す
	fmt.Printf("%s\n", s2)
}
