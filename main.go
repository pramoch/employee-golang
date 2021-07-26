package main

import (
	"context"
	DB "employee-golang/db"
	Employees "employee-golang/employees"
	Router "employee-golang/router"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	DB.Connect(ctx)

	defer func() {
		DB.Disconnect(ctx)
	}()

	Employees.Init()
	Router.Init()
}
