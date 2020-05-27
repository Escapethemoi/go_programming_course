package main

import "fmt"

type Paint struct { //anything on Paint
	X int
	Y int
	Z int
}

func main() {
	var s Paint //anything on Paint, but same as above
	s.X = 15
	s.Y = 23
	s.Z = -2
	fmt.Println("Määränpääsi on", s)
}
