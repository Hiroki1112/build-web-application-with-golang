//channels
package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _,v := range a {
		total += v
	}
	c <- total//send total to c...
}

func main() {
	a := []int{1,2,3,4,5,6,7}

	c := make(chan int)
	go sum(a[:len(a)/2],c)//from a[3] to a[6]
	go sum(a[len(a)/2:],c)//from a[0] to a[2]
	x,y := <-c, <-c

	fmt.Println(x,y,x+y)

}