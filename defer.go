package main

import (
	"fmt"
	"net/http"
)

// main
func main() {
	defer fmt.Pnrintln("world")
	test()
	fmt.Println("hello")
}

func test() {
	fmt.Println("world")
}
