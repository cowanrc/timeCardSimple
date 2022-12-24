package main

import (
	"log"
	"timeCardSimple/employee"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Printf("Creating your time card application")
	e := echo.New()

	e.File("/explorer", "ui/index.html")
	e.Static("/explorer", "ui")

	//routes
	e.POST("/employees", employee.CreateEmployeeHandler)
	e.GET("/employees", employee.GetAllEmployeeHandler)
	e.GET("/employees/:id", employee.GetEmployeeHandler)
	e.POST("/employees/ClockIn/:id", employee.ClockInHandler)
	e.POST("/employees/ClockOut/:id", employee.ClockOutHandler)
	e.DELETE("/employees/:id", employee.DeleteEmployeeHandler)

	log.Printf("listening on port 8080")
	e.Logger.Fatal((e.Start(":8080")))

}
