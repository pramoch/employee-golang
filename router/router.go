package Router

import (
	Employees "employee-golang/employees"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.GET("/employees", Employees.GetEmployees)

	r.Run("localhost:8080")
}
