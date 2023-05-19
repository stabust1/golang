package main

import "fmt"

var (
	n int    = 5
	m int    = 10
	c string = "abc"
)

func main() {
	var a int
	a = 10
	var i int = 11
	j := 12
	y := 13
	v := "Hello, World"
	fmt.Println(a)
	fmt.Println(i)
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(c)
	fmt.Println(j)
	fmt.Printf("%v, %T", y, y)
	fmt.Printf("%v, %T", v, v)
}
