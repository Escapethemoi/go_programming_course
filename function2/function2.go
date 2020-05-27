package main

import "fmt"

func sayHello(i int, greeting string) {
	fmt.Println(greeting, "Otto", i)
}

func main() {
	sayHello(90, "Muchos Nachos")
}
