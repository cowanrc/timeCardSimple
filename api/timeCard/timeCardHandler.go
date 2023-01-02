package timeCard

import (
	"net/http"
	"timeCardSimple/api"
	"timeCardSimple/database"

	"github.com/labstack/echo/v4"
)

func ClockInHandler(ctx echo.Context, isTesting bool) error {
	employeeId, idErr := api.GetEmployeeId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return nil
	}

	var employee database.TimeCard
	employee.EmployeeID = employeeId

	if isTesting {
		return nil
	}

	result, err := TimeCardService.ClockInEmployee(employee)
	if err != nil {
		ctx.JSON(err.Status, err)
		return nil
	}

	return ctx.JSON(http.StatusOK, result)
}

func ClockOutHandler(ctx echo.Context, isTesting bool) error {
	employeeId, idErr := api.GetEmployeeId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return nil
	}

	var employee database.TimeCard
	employee.EmployeeID = employeeId

	if isTesting {
		return nil
	}

	result, err := TimeCardService.ClockOutEmployee(employee)
	if err != nil {
		ctx.JSON(err.Status, err)
		return nil
	}

	return ctx.JSON(http.StatusOK, result)
}

func TotalTimeHandler(ctx echo.Context, isTesting bool) error {
	employeeId, idErr := api.GetEmployeeId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return nil
	}

	if isTesting {
		return nil
	}

	employee, getErr := TimeCardService.GetTotalTime(employeeId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return nil
	}

	return ctx.JSON(http.StatusOK, employee)
}
