package main

import "fmt"

func main() {
	x := 15
	a := &x // pointer to the memory address of x
	fmt.Println(a)
	fmt.Println(*a) // read through a memory address
	*a = 5
	fmt.Println(x)
	fmt.Println(*a)
	*a = *a**a
	fmt.Println(x)
}