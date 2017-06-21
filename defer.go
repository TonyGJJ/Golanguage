package main

import (
	"fmt"
)

// main
func main()  {
	defer fmt.Println("world")
	test()
	fmt.Println("hello")
}

func test() {
	fmt.Println("world")
}
