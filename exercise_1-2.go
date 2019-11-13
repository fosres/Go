//Echo2 prints its command line arguments

package main

import (
	"strconv"	
	"fmt"
	"os"
	)

func main() {
	s:=os.Args[0]+" "
	for i:=1; i < len(os.Args);i++ {
		s += "i:" + strconv.Itoa(i) + " " + os.Args[i] + " "
								
	}

	fmt.Println(s)
}
