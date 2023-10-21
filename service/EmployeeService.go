package service

import (
	"errors"

	"github.com/get2ashish/go-employee-api/models"
)

var employees = []models.Employee{
	{ID: "1", Name: "Ashish", Address: "Phoenix AZ", Salary: 2000},
	{ID: "2", Name: "Sachi", Address: "Phoenix AZ", Salary: 5000},
	{ID: "3", Name: "Avi", Address: "Phoenix AZ", Salary: 6000},
}

func GetAllEmployees() []models.Employee {
	return employees
}

func FindEmployeeById(id string) (*models.Employee, error) {
	for i, e := range employees {
		if e.ID == id {
			return &employees[i], nil
		}
	}
	return nil, errors.New("employee not found")
}

func CreateEmployee(employee models.Employee) {
	employees = append(employees, employee)
}
