package Employees

import "github.com/gin-gonic/gin"

type employee struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

var e []employee

func Init() {
	e = []employee{
		{Name: "Emma", Surname: "Cartner"},
		{Name: "Quinn", Surname: "Rivers"},
		{Name: "Amelia", Surname: "Burrows"},
	}
}

func GetEmployees(c *gin.Context) {
	c.JSON(200, e)
}
