package main

import "fmt"

type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{1, 2}  //has type vertex
	v2 = Vertex{x: 1}  //Y:0 is implicit
	v3 = Vertex{}      //X:0 and Y:0
	p  = &Vertex{1, 2} //has type of * vertex
)

func main() {
	//Struct
	v := Vertex{1, 2}
	v.x = 4
	fmt.Println(v.x)

	//Pointers To Struct
	z := Vertex{2, 3}
	p := &z
	p.x = 1e9
	fmt.Println(z)

	//Struct Literals
	fmt.Println(v1, p, v2, v3)
}
