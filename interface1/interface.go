package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rect struct {
	length float64
	width  float64
}
type circle struct {
	r float64
}

// created method area
func (r *rect) area() float64 {
	return r.length * r.width

}

// created method area of circle
func (c circle) area() float64 {
	return math.Pi * c.r * c.r

}

func main() {
	c1 := circle{6.9}
	r2 := rect{8, 7}
	shapes := []shape{&r2, c1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

}
