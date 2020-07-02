package main

import "fmt"

func main() {
	// 10個の要素を宣言します。要素の型はbyteの配列です。
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	
	// byteを含む２つのsliceを宣言します
	var a, b []byte

	// aポインタ配列の3つ目の要素から始まり、5つ目の要素まで
	a = ar[2:5]
	//現在aの持つ要素は：ar[2]、ar[3]とar[4]です。

	// bは配列arのもう一つのsliceです。
	b = ar[3:5]
	// bの要素は：ar[3]とar[4]です。
	println(len(a))
	println(cap(a))
	println(b)
	println("-----")
	//var numbers map[string]int
	// もうひとつのmapの宣言方法
	numbers := make(map[string]int)
	numbers["one"] = 1  //代入
	numbers["ten"] = 10 //代入
	numbers["three"] = 3
	fmt.Println("３つ目の数字は: ", numbers["three"]) // データの取得
	// "３つ目の数字は： 3"という風に出力されます。
}
