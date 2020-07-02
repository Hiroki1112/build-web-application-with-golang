package main

//import "fmt"

func add(a *int)int{
	*a = *a + 2 
	return *a
}

func main(){
	x:=3
	x1:=add(&x)
	println(x1)
}