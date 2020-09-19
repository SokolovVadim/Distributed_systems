package main

import "fmt"

func nil_ptr() {
	var p* int
	fmt.Println(p)
}

func play_with_ptrs() {
	i := 10 
	var p* int
	p = &i
	fmt.Printf("%d, %p\n", i, p)
	*p = 15
	j := 0
	p = &j
	*p = 10
	fmt.Printf("%d, %d, %p\n", i, j, p)
}

type Vertex struct {
	X float64
	Y float64
}

func try_struct() {
	v := Vertex{10.1, 12.1}
	v.X++
	v.Y--
	fmt.Println(v)

	// Ptr to struct

	p := &v
	p.X += 1e5
	fmt.Println(v)

	// Struct literals

	var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p0  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, v2, v3, p0)
}

func try_array() {
	
}

func main(){
	// nil_ptr()
	// play_with_ptrs()
	try_struct()
}
