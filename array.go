/*
数组
类型[n]T 是一个有n个类型为T的值的数组。
表达式

var a [10]int

定义变量a是一个有十个整数的数组。
数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是一个制约，但是请不要担心；Go 提供了更加便利的方式来使用数组。
*/

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var s [3]int
	fmt.Println(s[0])
	fmt.Println(s[len(s)-1])

	for i, v := range s {
		fmt.Printf("%d %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}
