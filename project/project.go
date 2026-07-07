package main

import (
	"fmt"
	"strings"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

var students []Student

func addStudent() {
	var student Student

	fmt.Print("Enter Student ID: ")
	fmt.Scanln(&student.ID)

	fmt.Print("Enter Student Name: ")
	fmt.Scanln(&student.Name)

	fmt.Print("Enter Student Age: ")
	fmt.Scanln(&student.Age)

	students = append(students, student)
	fmt.Println("\n Student added successfully!")
}

func viewStudents() {
	if len(students) == 0 {
		fmt.Println("\nNo students found.")
		return
	}

	fmt.Println("\n----- Student List -----")
	for _, student := range students {
		fmt.Printf("ID: %d | Name: %s | Age: %d\n",
			student.ID, student.Name, student.Age)
	}
}

func searchStudent() {
	var name string

	fmt.Print("Enter student name to search: ")
	fmt.Scanln(&name)

	for _, student := range students {
		if strings.EqualFold(student.Name, name) {
			fmt.Println("\nStudent Found")
			fmt.Printf("ID   : %d\n", student.ID)
			fmt.Printf("Name : %s\n", student.Name)
			fmt.Printf("Age  : %d\n", student.Age)
			return
		}
	}

	fmt.Println("\nStudent not found.")
}

func menu() {
	fmt.Println("\n===== Student Management System =====")
	fmt.Println("1. Add Student")
	fmt.Println("2. View Students")
	fmt.Println("3. Search Student")
	fmt.Println("4. Exit")
	fmt.Print("Enter your choice: ")
}

func main() {
	for {
		menu()

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addStudent()
		case 2:
			viewStudents()
		case 3:
			searchStudent()
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}



