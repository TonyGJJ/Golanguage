package main

import  "fmt"

func main(){
	var z[]int
	fmt.Println(z, len(z), cap(z))

	if z == nil {
		fmt.Println("nil")
	}

	var pow = []int{1, 2, 4, 8, 16}

	// 返回下标和值
	for i, v := range pow {
		fmt.Printf("2***%d = %d\n", i, v)
	}
}
