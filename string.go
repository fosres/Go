package main

import (
	"fmt"
)

func main() {
	a := "the"

	a = a + " swiss"

	b := []byte(a)

	b = append(b,0xa)

	fmt.Printf("%s\n",b)
}
