package main

import (
	"fmt"
	"os"
	"/home/teamvats/go-class/go-class/hello"
) 

func main() {
	if len(os.Args) > 1 {
		fmt.Println(hello.sayName(os.Args[1]))
	} else{
		fmt.Println(hello.sayName("world!"))
	}
}