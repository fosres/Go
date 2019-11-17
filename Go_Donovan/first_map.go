package main

import(
	"fmt"
)

func main() {
	
	url_map := make(map[string]uint8)

	url_map["https://en.wikipedia.org"] = 1

	url_map["https://wikia.com"] = 1

	fmt.Printf("Empty map: ",url_map)
}
