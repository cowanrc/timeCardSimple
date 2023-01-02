package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"timeCardSimple/api/employees"
	"timeCardSimple/api/timeCard"
	"timeCardSimple/database"

	"github.com/labstack/echo/v4"
)

const (
	is_testing = "is_testing"
)

func getenvBool(key string) (bool, error) {
	v, err := strconv.ParseBool(key)
	if err != nil {
		return false, err
	}
	return v, nil
}

func main() {
	key := os.Getenv(is_testing)
	isTesting, err := getenvBool(key)
	fmt.Println("!!!!!!!: ", isTesting)
	if err != nil {
		panic(err)
	}

	if !isTesting {
		fmt.Println("Creating DB")
		database.CreateDatabase()
	}

	log.Printf("Creating your time card application")
	e := echo.New()

	e.File("/swaggerui", "ui/index.html")
	e.Static("/swaggerui", "ui")

	//routes
	e.POST("/employees", func(c echo.Context) error {
		return employees.CreateEmployeeHandler(c, isTesting)
	})
	e.GET("/employees", func(c echo.Context) error {
		return employees.GetAllEmployeesHandler(c, isTesting)
	})

	e.GET("/employees/:id", func(c echo.Context) error {
		return employees.GetEmployeeHandler(c, isTesting)
	})

	e.DELETE("/employees/:id", func(c echo.Context) error {
		return employees.DeleteEmployeeHandler(c, isTesting)
	})

	e.PUT("/employees/ClockIn/:id", timeCard.ClockInHandler)
	e.PUT("/employees/ClockOut/:id", timeCard.ClockOutHandler)

	e.GET("/employees/TotalTime/:id", timeCard.TotalTimeHandler)

	log.Printf("listening on port 8080")
	e.Logger.Fatal((e.Start(":8080")))

}
