gpackage main

import "fmt"
import "flag"

func main() {
	var name string
	var times int
	flag.StringVar(&name, "name", "Anomuumi", "name to greet")
	flag.IntVar(&times, "times", 1, "times to greet")
	flag.Parse()

	for i:=0; i<times; i++{
		fmt.Println("Moikka", name)
	}
}
