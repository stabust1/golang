package main

import "fmt"

func main() {
	var a float32 = 3.14
	var b float32 = 3.14

	fmt.Printf("%v, %T\n", a+b, a+b)
	fmt.Printf("%v, %T\n", a-b, a-b)
	fmt.Printf("%v, %T\n", a*b, a*b)
	fmt.Printf("%v, %T\n", a/b, a/b)

}
