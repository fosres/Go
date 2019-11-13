package main

import(
	"fmt"
	"time"
)

func numbers(){
	for i:= 1; i <= 5; i++ {
		fmt.Printf("%d ",i)

	}
}

func alphabets(){
	for i:='a';i<='e';i++ {
		fmt.Printf("%c ",i)
	}
}

func main(){
	go numbers()
	go alphabets()
	time.Sleep(time.Millisecond*10)
	fmt.Println("main terminated")
	fmt.Println("swag")
}
