package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//Maps
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Range

	for i, j := range pow {
		fmt.Printf("2**%d = %d\n", i, j)
	}
	//Arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//Slices
	primes2 := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes2[1:4]
	fmt.Println(s)

	//Slices
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	x := names[0:2]
	y := names[1:3]
	fmt.Println(x, y)

	y[0] = "VVV"
	fmt.Println(x, y)
	fmt.Println(names)

	///////////////////////////
	g := []int{2, 3, 5, 7, 11, 13}
	g = g[1:4]
	fmt.Println(g)

	g = g[:2]
	fmt.Println(g)

	g = g[1:]
	fmt.Println(g)

	//////////////////
	m := []int{2, 3, 5, 7, 11, 13}
	printSlice(m)

	// Slice the slice to give it zero length.
	m = m[:0]
	printSlice(m)

	// Extend its length.
	m = m[:4]
	printSlice(m)

	// Drop its first two values.
	m = m[2:]
	printSlice(m)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
