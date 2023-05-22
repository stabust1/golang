package main

import "fmt"

type Color int

const (
	Red   Color = iota // 0
	Green              // 1
	Blue               // 2
)

func (c Color) String() string {
	names := [...]string{"Red", "Green", "Blue"}
	if c < Red || c > Blue {
		return "Unknown"
	}
	return names[c]
}

func main() {
	color := Red
	fmt.Println("Màu", color)
	color = Green
	fmt.Println("Màu", color)
	color = 5
	fmt.Println("Màu", color)
}
