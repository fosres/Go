// Fetch prints the content found at a URL.

package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"bytes"
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

func crawler(html_page []byte) [1024][1024] byte {
	
	search_str := []byte("https://")

	var urls [1024][1024]byte

	i := 0

	url_index := 0

	url_offset := 0	


	for (i < len(html_page)) && (url_index < 1024) {
		
		i = bytes.Index(html_page,search_str)
		
		for html_page[i] != 0x22 {
			
			urls[url_index][url_offset] = html_page[i]
			
			url_offset++

			i++
		}

		url_index++

		url_offset = 0

		i++

	}

	return urls
	
	
}

func main() {
	

//	i := 0 

	c := getPages(os.Args)

/*
	for c[i] != nil {
		fmt.Printf("%s\n",c[i])
		i++
	}
*/

	url_list := crawler(c[0])

	i := 0

	for url_list[0][i] != 0x0 {
		fmt.Printf("%c",url_list[0][i])
		i++
	}

	fmt.Printf("%c",0xa)


}
