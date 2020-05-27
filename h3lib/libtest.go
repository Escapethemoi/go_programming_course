package main

import "fmt"
import "bufio"
import "os"


//echo simulator: takes text input, loops it

func main() {
	fmt.Println("Enter text for echo:")
	reader := bufio.NewReader(os.Stdin)
	byteOption, _, _ := reader.ReadLine()
	echo := string(byteOption)
	for i := 0; i < 3; i++ {
	fmt.Println(echo)
	}
}
