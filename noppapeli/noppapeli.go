package main

import "fmt"
import "math/rand"
import "time"

func main() {
	aika := time.Now()
	kello := aika.Format(time.Kitchen)
	fmt.Println("Kello on", kello)
	noppa := (rand.Intn(20))
	fmt.Println("Heitit nopalla luvun", noppa)
}
