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

func extract_urls(html_page []byte) [1024][] byte {
	
	search_str := []byte("https://")

	var urls [1024][]byte

	i := 0

	url_index := 0
	
	len_html_page := len(html_page)

	for ( i < len_html_page && (url_index < 1024) ) {
		
		i = bytes.Index(html_page[i:],search_str)
		
		for html_page[i] != 0x22 {
			
			urls[url_index] = append(urls[url_index],html_page[i])
			
			i++
		}

		url_index++

		i++

	}

	return urls
	
	
}

func main() {
	

	var c [1024][]byte 
	
	c = getPages(os.Args)
	
	fmt.Printf("%s\n",c[0])

	url_list := extract_urls(c[0])

	i := 0

	for i < 1024 {
		
		fmt.Printf("%s\n",url_list[i])

		i++

	}



}
