package main

import "fmt"

//goでは関数も変数の一種
type testInt func(int) bool//関数の方を宣言

func isOdd(integer int) bool{
	if integer%2==0{
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int{
	var result []int
	for _, value := range slice {
		//f()でキャスト
		//1個ずつ値を取り出して、f(関数)で返り値を受け取るtrueならばif文内が実行される
		if f(value){
			result = append(result, value)
		}
	}
	return result
}

func main(){
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)    // 関数の値渡し
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)  // 関数の値渡し
	fmt.Println("Even elements of slice are: ", even)
}