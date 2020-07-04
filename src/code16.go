package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)//close channel
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)//cap:get size of channel. here, cap(c) is 10.
	for i := range c {
		fmt.Println(i)
	}
}