package main

import (
	Employees "employee-golang/employees"

	"github.com/gin-gonic/gin"
)

func main() {
	Employees.Init()

	r := gin.Default()
	r.GET("/employees", Employees.GetEmployees)

	r.Run("localhost:8080")
}
