package main

import (
	"errors"
	"fmt"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

const (
	HR    = "HR"
	IT    = "IT"
	Sales = "Sales"
)

var Employees []Employee

func AddEmployee(id int, name string, age int, department string) error {
	for _, emp := range Employees {
		if emp.ID == id {
			return errors.New("Employee with same ID already exists")
		}
	}

	if age < 18 {
		return errors.New("Employee age should be greater than 18")
	}

	employee := Employee{ID: id, Name: name, Age: age, Department: department}
	Employees = append(Employees, employee)
	return nil
}

func SearchEmployee(query any) (*Employee, error) {

	switch query.(type) {
	case int:
		for _, emp := range Employees {
			if emp.ID == query {
				return &emp, nil
			}
		}
		return nil, errors.New("Employee not found")
	case string:
		queryName, ok := query.(string)
		if !ok {
			return nil, errors.New("Invalid search query")
		}
		for _, emp := range Employees {
			if strings.EqualFold(emp.Name, queryName) {
				return &emp, nil
			}
		}
		return nil, errors.New("Employee not found")
	}
	return nil, errors.New("Invalid search query")
}

func ListEmployeesByDept(department string) []Employee {
	var empList []Employee
	for _, emp := range Employees {
		if emp.Department == department {
			empList = append(empList, emp)
		}
	}
	return empList
}

func CountEmployeesByDept(department string) int {
	count := 0
	for _, emp := range Employees {
		if emp.Department == department {
			count++
		}
	}
	return count
}

func displayEmployees(e *Employee) {
	fmt.Println("====================================")
	fmt.Println("ID: ", e.ID)
	fmt.Println("Name: ", e.Name)
	fmt.Println("Age: ", e.Age)
	fmt.Println("Department: ", e.Department)
	fmt.Println("====================================")
}

// func main() {

// 	//Adding Employees
// 	err := AddEmployee(1, "John", 25, HR)
// 	if err != nil {
// 		fmt.Println("Error adding employee: ", err)
// 	}

// 	err = AddEmployee(2, "Jane", 30, IT)

// 	if err != nil {
// 		fmt.Println("Error adding employee: ", err)
// 	}

// 	err = AddEmployee(3, "Doe", 17, Sales)
// 	if err != nil {
// 		fmt.Println("Error adding employee: ", err)
// 	}

// 	err = AddEmployee(4, "Jack", 25, HR)
// 	if err != nil {
// 		fmt.Println("Error adding employee: ", err)
// 	}
// 	//Searching Employees by ID
// 	emp, err := SearchEmployee(1)
// 	if err != nil {
// 		fmt.Println("Error searching employee: ", err)
// 	} else {
// 		fmt.Println("Employee found")
// 		displayEmployees(emp)
// 	}

// 	//Searching Employees by Name
// 	emp, err = SearchEmployee("Jane")
// 	if err != nil {
// 		fmt.Println("Error searching employee: ", err)
// 	} else {
// 		fmt.Println("Employee found")
// 		displayEmployees(emp)
// 	}

// 	//List Employees by Department
// 	empList := ListEmployeesByDept(HR)
// 	fmt.Println("Employees in HR: ")
// 	for _, emp := range empList {
// 		displayEmployees(&emp)
// 	}

// 	//Count Employees by Department
// 	count := CountEmployeesByDept(HR)
// 	fmt.Println("Count of Employees in HR: ", count)
// }
