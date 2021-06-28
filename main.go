package main

import (
	Employees "employee-golang/employees"
	Router "employee-golang/router"
)

func main() {
	Employees.Init()
	Router.Init()
}
