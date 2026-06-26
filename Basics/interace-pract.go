package main

import "fmt"

type area interface {
	rectarea() float32
}

type rectangle struct {
	length  float32
	breadth float32
}

func (r rectangle) rectarea() float32 {
	return r.length * r.breadth
}

func main() {

	r2 := rectangle{3, 4}
	fmt.Println(r2.rectarea())

}
