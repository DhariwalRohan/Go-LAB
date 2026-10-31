// package main

// import "fmt"

// func modifyMarks(marks *float64) {
// 	*marks = *marks + 5
// }

// type Student struct {
// 	Name  string
// 	Age   int
// 	Marks float64
// }

// func main() {
// 	// 1. Pointer manipulation
// 	age := 20

// 	fmt.Println("Age:", age)
// 	fmt.Println("Address of Age:", &age)
// 	fmt.Println("Age using pointer:", *(&age))

// 	// 2. Pass-by-reference
// 	fmt.Println("\nBefore modifying marks:", 85.5)
// 	marks := 85.5

// 	modifyMarks(&marks)

// 	fmt.Println("After modifying marks:", marks)

// 	// 3. Allocate struct using new()
// 	student := new(Student)

// 	student.Name = "Rahul"
// 	student.Age = 20
// 	student.Marks = 85.5

// 	fmt.Println("\nStudent Details:")
// 	fmt.Println("Name:", student.Name)
// 	fmt.Println("Age:", student.Age)
// 	fmt.Println("Marks:", student.Marks)

// 	// Modify fields through pointer
// 	student.Age = 21
// 	student.Marks = 90.5

// 	fmt.Println("\nModified Student Details:")
// 	fmt.Println("Name:", student.Name)
// 	fmt.Println("Age:", student.Age)
// 	fmt.Println("Marks:", student.Marks)
// }
