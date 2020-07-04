//interface:
//https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/02.6.md

package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age int
	phone string
}

func (h Human) String() string{
	return "Name:" + h.name + ",Age:" + strconv.Itoa(h.age) + "years old, Contact:" + h.phone
}

func main(){
	Bob := Human{"Mike",100,"xxx-3939-0334"}
	fmt.Println("This human is :",Bob)
}