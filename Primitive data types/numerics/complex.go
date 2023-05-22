package main

import "fmt"

func main() {
	var a complex64 = 1 + 2i
	var b complex128 = complex(10, 20)
	fmt.Printf("%v, %T\n", real(a), real(a))
	fmt.Printf("%v, %T\n", imag(a), imag(a))

	fmt.Printf("%v, %T\n", real(b), real(b))
	fmt.Printf("%v, %T\n", real(b), real(b))

}
