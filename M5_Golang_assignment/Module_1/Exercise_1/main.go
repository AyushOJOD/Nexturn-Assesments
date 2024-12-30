package main

import (
	"errors"
	"fmt"
)

// Employee struct for employee details
type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

// Array to store employees
var employees []Employee

// AddEmployee adds a new employee after validation
func AddEmployee(id int, name string, age int, department string) error {
	// Validate ID uniqueness
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("ID must be unique")
		}
	}

	// Validate age
	if age <= 18 {
		return errors.New("Age must be greater than 18")
	}

	// Add employee to the list
	employees = append(employees, Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})

	return nil
}

func main() {
	fmt.Println("Employee Management System")

	// Test adding employees
	if err := AddEmployee(1, "John Doe", 25, "IT"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee added successfully")
	}

	if err := AddEmployee(1, "Jane Smith", 22, "HR"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee added successfully")
	}

	if err := AddEmployee(2, "Alice Brown", 17, "Finance"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee added successfully")
	}

	fmt.Println("Current Employees:", employees)
}
