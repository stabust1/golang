package main

import "fmt"

func main() {
	var a bool
	var i bool
	var j bool
	a = 10 == 20
	i = 10 <= 20
	j = 10 >= 20
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
}
