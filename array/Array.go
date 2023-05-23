package main

import "fmt"

func main() {

	countries := [5]string{"one", "two", "three", "four", "five"}

	slice := countries[1:4]

	fmt.Println(slice)

	slice[0] = "six"
	fmt.Println(countries)
}
