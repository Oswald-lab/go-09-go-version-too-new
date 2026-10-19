package main

import "fmt"

func main() {
	for i := range 3 { // range-over-int needs Go 1.22+
		fmt.Println("tick", i)
	}
}
