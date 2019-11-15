// Fetch prints the content found at a URL.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


func getPages(a []string) [1024][]byte {

	i := 0 

	var ans [1024][]byte

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
			os.Exit(1)
		}

		ans[i] = b

		i++
		
	}
	
	return ans
}
func main() {
	

	i := 0 

/*
	var ans [1024][]byte

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
			os.Exit(1)
		}

		ans[i] = b

		i++
		
	}
	
	i = 0
*/
	c := getPages(os.Args)

	for c[i] != nil {
		fmt.Printf("%s\n",c[i])
		i++
	}

}
