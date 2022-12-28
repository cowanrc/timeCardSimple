package employee

import (
	"log"
	"net/http"
	"strconv"
	"timeCardSimple/database"
	"timeCardSimple/errors"

	"github.com/labstack/echo/v4"
)

type NewEmployee struct {
	Name        string `json:"Name,omitempty"`
	EmployeeID  int    `json:"employeeID,omitempty"`
	DateOfBirth string `json:"DoB,omitempty"`
}

type Employee struct {
	Name        string `json:"name,omitempty"`
	EmployeeID  int    `json:"employeeID,omitempty"`
	ClockIn     string `json:"clockIn,omitempty"`
	ClockOut    string `json:"clockOut,omitempty"`
	TotalTime   string `json:"totalTime,omitempty"`
	DateOfBirth string `json:"dob,omitempty"`
}

var TimeCard = make(map[int]*Employee)

// CreateEmployeeHandler to enter name and DOB and get an employee ID in return
func CreateEmployeeHandler(ctx echo.Context) error {
	// var newEmployee NewEmployee
	// newEmployee.EmployeeID = seq

	// err := ctx.Bind(&newEmployee)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, "Error binding the structure")
	// }

	// log.Printf("Employee Name is : %s", newEmployee.Name)
	// log.Printf("Employee id is: %v", newEmployee.EmployeeID)

	// seq++

	// m := map[string]string{
	// 	"name":       newEmployee.Name,
	// 	"employeeID": strconv.Itoa(newEmployee.EmployeeID),
	// }

	// var employee Employee

	// employee.Name = newEmployee.Name
	// employee.EmployeeID = newEmployee.EmployeeID
	// employee.DateOfBirth = newEmployee.DateOfBirth
	// TimeCard[newEmployee.EmployeeID] = &employee

	// return ctx.JSON(http.StatusAccepted, m)
	var employee database.Employee

	if err := ctx.Bind(&employee); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return err
	}

	result, saveErr := EmployeeService.CreateEmployee(employee)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return nil
	}

	return ctx.JSON(http.StatusCreated, result)

}

func GetAllEmployeesHandler(ctx echo.Context) error {
	employees, getErr := EmployeeService.GetAllEmployees()
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return nil
	}

	return ctx.JSON(http.StatusOK, employees)
}

func GetEmployeeHandler(ctx echo.Context) error {
	employeeId, idErr := getEmployeeId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return echo.ErrNotFound
	}

	employee, getErr := EmployeeService.GetEmployee(employeeId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return nil
	}

	return ctx.JSON(http.StatusOK, employee)

}

func DeleteEmployeeHandler(ctx echo.Context) error {

	employeeId, idErr := getEmployeeId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return echo.ErrNotFound
	}

	if err := EmployeeService.DeleteEmployee(employeeId); err != nil {
		ctx.JSON(err.Status, err)
		return echo.ErrBadRequest
	}

	return ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func ClockInHandler(ctx echo.Context) error {
	employeeId, idErr := getEmployeeId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return echo.ErrNotFound
	}

	var employee database.Employee
	employee.EmployeeID = employeeId

	result, err := EmployeeService.ClockInEmployee(employee)
	if err != nil {
		ctx.JSON(err.Status, err)
		return echo.ErrBadRequest
	}

	return ctx.JSON(http.StatusOK, result)
}

func ClockOutHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	employee := TimeCard[id]

	if err != nil {
		log.Printf("User must provide integer")
		return echo.NewHTTPError(http.StatusBadRequest, "User must provide an integer for an ID")
	}

	if !employeeExists(id) {
		log.Printf("Employee ID: %d does not exist in our system. Either this employee has been removed or has yet to be added.", id)
		return echo.NewHTTPError(http.StatusNotFound, "Employee ID: "+strconv.Itoa(id)+" does not exist in our system. Either this employee has been removed or has yet to be added.")
	}

	if employee.ClockIn == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User must first clock in before they can clock out.")
	} else if employee.ClockOut != "" {
		return echo.NewHTTPError(http.StatusBadRequest, "You cannot clock out multiple times. Without Clocking in again first")
	}

	employeeClockOut(id)

	m := map[string]string{
		"name":       employee.Name,
		"employeeID": strconv.Itoa(id),
		"clockIn":    employee.ClockIn,
		"clockOut":   employee.ClockOut,
	}

	return ctx.JSON(http.StatusAccepted, m)
}
