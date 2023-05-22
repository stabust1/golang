package main

import "fmt"

func main() {
	var a int8 = 5
	var b int8 = 2

	fmt.Printf("%v, %T\n", a+b, a+b)
	fmt.Printf("%v, %T\n", a+b, a+b)
	fmt.Printf("%v, %T\n", a*b, a*b)
	fmt.Printf("%v, %T\n", a/b, a/b)

	fmt.Printf("%v, %T\n", a&b, a&b)
	fmt.Printf("%v, %T\n", a|b, a|b)
	fmt.Printf("%v, %T\n", a^b, a^b)
	fmt.Printf("%v, %T\n", a&^b, a&^b)

	fmt.Printf("%v, %T\n", a<<b, a<<b)
	fmt.Printf("%v, %T\n", a>>b, a>>b)
}
