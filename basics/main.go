package main

import (
	"fmt"
	"math"
	"runtime"
	// "math/rand"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func intToString(i int) (s string) {
	s = fmt.Sprintf("%d", i)
	return
}

func floatToInt(f float64) (i int) {
	i = int(f)
	return
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func loopAndPrint() {
	for i := 0; i < 10; i++ {
		if i >= 3 {
			fmt.Printf("Hi, this is the %dth loop\n", i + 1)
		} else {
			switch i {
				case 0: {
					fmt.Printf("Hi, this is the %dst loop\n", i + 1)
				}
				case 1: {
					fmt.Printf("Hi, this is the %dnd loop\n", i + 1)
				}
				case 2: {
					fmt.Printf("Hi, this is the %drd loop\n", i + 1)
				}
			}
		}
	}
}

func sqrt(x float64) float64 {
	z := 1.0
	
	for math.Abs(x - math.Pow(z, 2)) > 0.1 {
		fmt.Printf("%f\n", x - z)
		z -= (z * z - x) / (2 * z)
	}
	
	z = math.Round(z)

	return z
}

func findOS() {
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Printf("%s darwin!", os)
		
		case "linux":
			fmt.Printf("%s linux!", os)

		default:
			fmt.Printf("%s", os)
	}
}

func deferredFunc() {
	defer fmt.Printf(" world!")
	
	fmt.Printf("Hello")
}

func pointers() {
	i := 25

	p := &i

	*p = 24
	fmt.Printf("%d\n", i)
	*p = *p / 6
	fmt.Printf("%d\n", i)
}

type Vertex struct {
	X int
	Y int
}

func types() {
	a := Vertex{1, 2}
	b := Vertex{X: 12, Y: 13}
	c := &Vertex{1, 2}
	
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", *c)
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	
	b := [6]int{1, 2, 3, 4, 5, 6}
	
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
}

func slices() {
	primes := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	
	var s []int = primes[1:5]
	s[0] = 3
	fmt.Println(s)
	fmt.Println(primes)
	
	q := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Println(q)
	
	vxs := []Vertex{
		{1, 2},
		{X: 3, Y: 4},
		{5, 6},
	}
	fmt.Println(vxs)
	
	def := primes[:]
	defTwo := primes[:3]
	defThree := primes[1:]
	fmt.Println(def, defTwo, defThree)
	
	fmt.Println(cap(def), len(def), cap(primes), len(primes))
	
	var n []int
	fmt.Println(n, len(n), cap(n))
	
	if n == nil {
		fmt.Println("n slice is nil!")
	}
}

func dynamicSlices() {
	// This is a dynamically made array
	a := make([]int, 5)
	fmt.Println(len(a), cap(a))
	
	// This is how you append elements to a slice
	a = append(a, 6)
	fmt.Println(len(a), cap(a))
	
	// This is a regular capped array
	b := make([]int, 3, 5)
	fmt.Println(len(b), cap(b))
	
	c := b[:cap(b)]
	fmt.Println(len(c), cap(c))

	d := b[1:]
	fmt.Println(len(d), cap(d))
}

func ranges() {
	var pow = []int{2, 4, 8, 16, 32, 64, 128, 256}
	
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i + 1, v)
	}
	
	// All different assignment combinations for this
	// for _, v := range pow
	// for i, _ := range pow
	// for i := range pow
}

func main() {
	// var limit int = 31
	// var randomNumber int = rand.Intn(limit)
	// const hello = "Hello World!"
	
	// var s = intToString(123)
	// var i = floatToInt(123.693)
	
	// x, y := split(20)
	
	// var powered float64 = pow(2, 2, 5)

	// var sqrtOfFour float64 = sqrt(16.0)
	
	// loopAndPrint()
	// findOS()
	// deferredFunc()
	// pointers()
	// types()
	// arrays()
	// slices()
	// dynamicSlices()
	ranges()
	// fmt.Printf("Square root of 16 is: %f\n", sqrtOfFour)
	// fmt.Printf("2 to the power of 2 is: %f\n", powered)
	// fmt.Printf("Converted: %s, %d to types %T, %T\n", s, i, s, i)
	// fmt.Printf("Numbers are: %d, %d\n", x, y)
	// fmt.Println("My favourite number is: ", randomNumber)
	// fmt.Println(hello)
}