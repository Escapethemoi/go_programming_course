package main

import "fmt"

func main() {
	for i:=0; i<10; i++ {
		if i==0 || i==2 || i==4 || i==6 || i==8  {
			fmt.Println("#=#=#=#=#=")
		} else {
			fmt.Println("=#=#=#=#=#")
		}
	}
}
