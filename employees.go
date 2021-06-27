package main

import "github.com/gin-gonic/gin"

type employee struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type employees []employee

func newEmployees() employees {
	employees := []employee{
		{Name: "Emma", Surname: "Cartner"},
		{Name: "Quinn", Surname: "Rivers"},
		{Name: "Amelia", Surname: "Burrows"},
	}
	return employees
}

func (e employees) getEmployees(c *gin.Context) {
	c.JSON(200, e)
}
