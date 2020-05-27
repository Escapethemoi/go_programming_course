package main

import "fmt"

func main () {
	var planeetat = []string{"Saturnussi", "Jupitero", "Venus", "Aanus", "Sulttaanin tähtisumu"}

	for _, planeetta := range planeetat {
		fmt.Println(planeetta)
	}
}
