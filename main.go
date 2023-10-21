package main

import (
	"net/http"

	"github.com/get2ashish/go-employee-api/models"
	"github.com/get2ashish/go-employee-api/service"
	"github.com/gin-gonic/gin"
)

func getAllEmployees(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, service.GetAllEmployees())
}

func getEmployeeById(context *gin.Context) {
	id := context.Param("id")
	employee, error := service.FindEmployeeById(id)
	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, employee)
}

func createEmployee(context *gin.Context) {
	var payload models.Employee
	error := context.BindJSON(&payload)

	if error != nil {
		return
	}

	service.CreateEmployee(payload)
	context.IndentedJSON(http.StatusCreated, payload)
}

func main() {
	router := gin.Default()
	router.GET("/employee", getAllEmployees)
	router.GET("/employee/:id", getEmployeeById)
	router.POST("/employee", createEmployee)
	router.Run(":8080")
}
