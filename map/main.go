package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func main() {
	// Khai báo map với key là string và value là struct Student
	students := make(map[string]Student)

	// Add student information to the map
	students["Alice"] = Student{Name: "Toan", Age: 23, Grade: 2}
	students["Bob"] = Student{Name: "Truong", Age: 23, Grade: 2}
	students["Charlie"] = Student{Name: "Duong", Age: 23, Grade: 3}

	// Access student information in the map
	Toan := students["Toan"]
	fmt.Println("Name:", Toan.Name)
	fmt.Println("Age:", Toan.Age)
	fmt.Println("Grade:", Toan.Grade)

	// Check the existence of students in the map
	if student, ok := students["Dave"]; ok {
		fmt.Println("Name:", student.Name)
		fmt.Println("Age:", student.Age)
		fmt.Println("Grade:", student.Grade)
	} else {
		fmt.Println("Dave not found")
	}
	//Browse all students in the map
	for _, student := range students {
		fmt.Println("Name:", student.Name)
		fmt.Println("Age:", student.Age)
		fmt.Println("Grade:", student.Grade)
		fmt.Println("-----")
	}
}
