package main

import(
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"math"
)

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

func slice_to_string(values[] int) string {
	valuesText := []string{}

    // Create a string slice using strconv.Itoa.
    // ... Append strings to it.
    for i := range values {
        number := values[i]
        text := strconv.Itoa(number)
        valuesText = append(valuesText, text)
    }

    // Join our string slice.
    result := strings.Join(valuesText, ", ")
    return result
}

func try_array() {
	var array[2] string
	array[0] = "Hello"
	array[1	] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	primes := [6] int {2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var slice[] int = primes[1:4]
	slice_str := slice_to_string(slice)
	fmt.Println(slice_str)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// slice literals 

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// slice defaults

	slice0 := []int{2, 3, 5, 7, 11, 13}

	slice0 = slice0[1:4]
	fmt.Println(slice0)

	slice0 = slice0[:2]
	fmt.Println(slice0)

	slice0 = slice0[1:]
	fmt.Println(slice0)
}

func printSlice(slice []int) {
	fmt.Printf("len = %d, cap = %d\nslice: %v\n", len(slice), cap(slice), slice)
}

func try_more_slices() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func nil_slice() {
	var slice[] int
	fmt.Println(slice, len(slice), cap(slice))
	if slice == nil {
		fmt.Println("nil!")
	}
}

func printSliceStr(str string, slice []int) {
	fmt.Printf("%s: len = %d, cap = %d\nslice: %v\n", str, len(slice), cap(slice), slice)
}

func create_slice_with_make() {
	a := make([]int, 5)
	printSliceStr("a", a)

	b := make([]int, 0, 5)
	printSliceStr("b", b)

	c := b[:2]
	printSliceStr("c", c)

	d := c[2:5]
	printSliceStr("d", d)
}

func slices_of_slices() {
	// Create a tic tac toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// Players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i:= 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func append_to_slice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func try_range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func skip_index() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func Pic(dx, dy int) [][]uint8 {
	var matrix = make([]([]uint8), dy)
	// slice := make([]uint8, 5)
	for i := 0; i < dx; i++ {
		matrix[i] = make([]uint8, dx)
		for j := 0; j < dy; j++ {
			matrix[i][j] = uint8((i + j) / 2)	
		}
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			fmt.Printf("%c ", matrix[i][j])		
		}
		fmt.Printf("\n")
	}
	return matrix
}

func try_map() {
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func map_literals() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func top_level() {
	type Vertex struct {
	Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func mutating_maps() {
	m := make(map[string] int)

	m["Height"] = 190
	m["Weight"] = 87
	fmt.Println("Map: ", m)
	fmt.Println("The value: ", m["Height"])

	delete(m, "Weight")
	fmt.Println(m)

	v, ok := m["Height"]
	fmt.Println("The value:", v, "Present?", ok)
	v, ok = m["Weight"]
	fmt.Println("The value:", v, "Present?", ok)
}

func isWord(c rune) bool{
	return !unicode.IsLetter(c) 
}

func WordCount(str string) map[string]int {
	m := make(map[string] int)
	str_slices := strings.Fields(str)
	// fmt.Println(str_slices)
	for _, v := range str_slices {
		m[v]++
	}

	return m
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func function_values() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first + second
		return ret
	}
}

func run_fibonacci(n int) {
	f := fibonacci()
	for i := 0; i < n; i++ {
		fmt.Println(f())
	}
}

func main(){
	// nil_ptr()
	// play_with_ptrs()
	// try_struct()
	// try_array()
	// try_more_slices()
	// nil_slice()
	// create_slice_with_make()
	// slices_of_slices()
	// append_to_slice()
	// try_range()
	// skip_index()
	// Pic(10, 10)
	// try_map()
	// top_level()
	// mutating_maps()
	// fmt.Println(WordCount("I learn Go & sleep well"))
	// function_values()
	// closures()
	run_fibonacci(10)
}
