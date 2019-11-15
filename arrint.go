package main

import(
	"fmt"
)

func main() {
	
	a := [3]int{12,78,50} // short hand declaration to create array
	
	var b []int = a[:1] //creates a slice for a[1] to a[3]e

	fmt.Println(b)

}
