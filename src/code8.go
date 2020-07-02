package main
import "fmt"

type Car struct{
	name string
	tank int
	country string
}

//匿名フィールド
type Cars struct {
	Car
	hoge string
}

func main(){
	honda := Cars{Car{"CR-V",30,"japan"},"good"}
	fmt.Printf("Name:%s, tank:%d, made in %s\n",honda.name,honda.tank,honda.country)
}