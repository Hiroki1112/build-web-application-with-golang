package main

import (
	"fmt"
	"math"
)

type Recrangle struct {
	width,height float64
}

type Circle struct {
	radius float64
}

type Human struct{
	weight,height float64
}

func (r Recrangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64{
	return c.radius * c.radius * math.Pi
}

func (h Human) BMI() float64{
	return h.weight/(h.height * h.height)
}

func main(){
	r1 := Recrangle{12, 2}
	r2 := Recrangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	h1 := Human{56,1.66}
	h2 := Human{76,1.80}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
	fmt.Println("BMI of h1 is: ", h1.BMI())
	fmt.Println("BMI of h2 is: ", h2.BMI())
}

