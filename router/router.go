package Router

import (
	Branches "employee-golang/branches"
	Employees "employee-golang/employees"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.GET("/employees/:id", Employees.GetEmployeeById)
	r.GET("/employees", Employees.GetEmployees)
	r.GET("/positions", Employees.GetPositions)
	r.GET("/branches/:id", Branches.GetBranchById)
	r.GET("/branches", Branches.GetBranches)

	r.Run("localhost:8080")
}
