package main

import "fmt"

func area(x int) int {
	return x*x
}

func main() {
	side := 5
	a := area(side)
	fmt.Println(a, "m2")
	fmt.Printf("%v m2\n", a)
}
