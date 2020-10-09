package main

import "fmt"

var num1, num2 float64 = 5.6, 9.5

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "there"

	fmt.Println(multiple(w1, w2))

	var a int = 62
	var b float64 = float64(a)

	x := a

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", b)
}