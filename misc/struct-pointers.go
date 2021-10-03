package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println("v =", v)
	p := &v
	fmt.Println("p =", p)
	fmt.Println("*p =", *p)
	p.X = 1e9
	(*p).Y = 3
	fmt.Println("v =", v)
}

