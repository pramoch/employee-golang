package main

import "github.com/gin-gonic/gin"

func main() {
	employees := newEmployees()
	r := gin.Default()

	r.GET("/employees", employees.getEmployees)

	r.Run("localhost:8080")
}
