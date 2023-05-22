package main

import "fmt"

func main() {
	var a string = "Hello World"
	var b []byte = []byte(a)

	fmt.Printf("%v, %T\n", b, b)

}
