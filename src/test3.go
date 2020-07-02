package main

import (
	"fmt"
	//"os"
)

//グループ化による宣言
/*
const (
	hoge = 1000
	fuga = 3.14
	hoho = 1.4142
)
*/

const (
	aa int = iota
	bb
	cc
)

/*
var (
	i int
	pi float64
	prefix string
)
*/
//戦闘が大文字で始まる変数/関数はエクスポート可能
//小文字の場合はエクスポートできない(プライベート変数)

func main() {
	var arr [10]int                                 // int型の配列を宣言します。
	a := [10]int{1,2,3}
	b := [...]int{1,2,3,4,5}
	c := [2][2]int{{1,2,3},{4,5,6}}
	//動的な配列
	var fslice []int
	slice := []byte{1,2,3}
	arr[0] = 42                                     // 配列のインデックスは0からはじまります。
	arr[1] = 13                                     // 代入操作
	fmt.Printf("The first element is %d\n", arr[0]) // データを取得して、42を返します。
	fmt.Printf("The last element is %d\n", arr[9])  //値が代入されていない最後の要素を返します。デフォルトでは0が返ります。
	fmt.Println(aa, bb, cc)
}
