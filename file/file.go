package main

import "fmt"
import "io/ioutil"

func readFile() {
	bytes, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	fmt.Println(string(bytes))
}

func main() {
	readFile()
}
