package employees

import (
	"net/http"
	"timeCardSimple/api"
	"timeCardSimple/database"
	"timeCardSimple/errors"

	"github.com/labstack/echo/v4"
)

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
	employeeId, idErr := api.GetEmployeeId(ctx.Param("id"))
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

	employeeId, idErr := api.GetEmployeeId(ctx.Param("id"))
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
