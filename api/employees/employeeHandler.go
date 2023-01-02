package employees

import (
	"net/http"
	"timeCardSimple/api"
	"timeCardSimple/database"
	"timeCardSimple/errors"

	"github.com/labstack/echo/v4"
)

func CreateEmployeeHandler(ctx echo.Context, isTesting bool) error {
	var employee database.Employee

	if err := ctx.Bind(&employee); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return err
	}

	if isTesting {
		return nil
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
		return nil
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
		return nil
	}

	if err := EmployeeService.DeleteEmployee(employeeId); err != nil {
		ctx.JSON(err.Status, err)
		return nil
	}

	return ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
