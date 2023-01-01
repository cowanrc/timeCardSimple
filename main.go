package main

import (
	"log"
	"timeCardSimple/api/employees"
	"timeCardSimple/api/timeCard"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Printf("Creating your time card application")
	e := echo.New()

	// e.File("/explorer", "ui/index.html")
	// e.Static("/explorer", "ui")
	e.Static("/swaggerui", "cmd/api/swaggerui")

	//routes
	e.POST("/employees", employees.CreateEmployeeHandler)
	e.GET("/employees", employees.GetAllEmployeesHandler)
	e.GET("/employees/:id", employees.GetEmployeeHandler)
	e.DELETE("/employees/:id", employees.DeleteEmployeeHandler)

	e.PUT("/employees/ClockIn/:id", timeCard.ClockInHandler)
	e.PUT("/employees/ClockOut/:id", timeCard.ClockOutHandler)

	e.GET("/employees/TotalTime/:id", timeCard.TotalTimeHandler)

	log.Printf("listening on port 8080")
	e.Logger.Fatal((e.Start(":8080")))

}
