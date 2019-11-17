package main

import(
	"fmt"
	"net"
)

func main() {
	
	addr, err := net.LookupIP("https://donate.wikimedia.org/wiki/Special:FundraiserRedirector?utm_source=donate&amp;utm_medium=sidebar&amp;utm_campaign=C13_en.wikipedia.org&amp;uselang=en")
	
	if err != nil {
		fmt.Println("unknown error :",err)
	} else{
		fmt.Println("IP address of website is: ",addr)
	}


}
