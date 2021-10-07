package Router

import (
	Branches "employee-golang/branches"
	Employees "employee-golang/employees"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	// Employees
	r.GET("/employees/:id", Employees.GetEmployeeById)
	r.GET("/employees", Employees.GetEmployees)

	// Positions
	r.GET("/positions", Employees.GetPositions)

	// Branches
	r.GET("/branches/:id", Branches.GetBranchById)
	r.GET("/branches", Branches.GetBranches)
	r.POST("/branches", Branches.AddBranch)

	r.Run("localhost:8080")
}
