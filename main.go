package main

import (
	"github.com/gin-gonic/gin"
)

type employee struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	employees := []employee{
		{Name: "Emma", Surname: "Cartner"},
		{Name: "Quinn", Surname: "Rivers"},
	}

	r := gin.Default()
	r.GET("/employees", func(c *gin.Context) {
		c.JSON(200, employees)
	})

	r.Run("localhost:8080")
}
