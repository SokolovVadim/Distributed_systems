package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func print(max_val int){
	for i := 0; i < max_val; i++ {
		fmt.Println(i)
	}
}

func while_loop(val, max_val int){
	for max_val < max_val {
		val++;
	}
	fmt.Println(val);
}

func infinite_loop() {
	for {
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("result (%g) is supposed to be less than lim(%g)\n", v, lim)
	}
	return lim
}

// prev := 0.0; z != prev; prev = z
func Sqrt(x float64) float64 {
	z := x / 2
	for z * z != x {
		z -= (z * z - x) / (2 * z)
		fmt.Printf("%g ", z)
	}
	return z
}

func os_switch() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	default:
		// freebsd, openbsd,
		// plan9, windows...
		return os
	}
}

func when_is_Saturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func greating() {
	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("Good morning!")
	case t < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func defer_performing() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func stacking_defers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main(){
	// print(5);
	// while_loop(5, 10)
	// infinite_loop()
	// fmt.Println(sqrt(1024), sqrt(-4), sqrt(2))
	// fmt.Println(pow(2, 11, 1024))
	// Sqrt(100)
	// fmt.Println(os_switch())
	// when_is_Saturday()
	// greating()
	// defer_performing()
	stacking_defers()
}
