package main

import "fmt"

func reply() {
	fmt.Println("Test")
}

func main() {
	extraStep()
}

func extraStep() {
	reply()
}
