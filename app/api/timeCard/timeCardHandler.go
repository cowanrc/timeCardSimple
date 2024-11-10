package timeCard

// import (
// 	"net/http"
// 	"timeCardSimple/api"
// 	"timeCardSimple/database"

// 	"github.com/labstack/echo/v4"
// )

// func ClockInHandler(ctx echo.Context, isTesting bool) error {
// 	employeeId, idErr := api.GetEmployeeId(ctx.Param("id"))
// 	if idErr != nil {
// 		ctx.JSON(idErr.Status, idErr)
// 		return nil
// 	}

// 	var employee database.TimeCard
// 	employee.EmployeeID = employeeId

// 	//improper use of unit test handling
// 	if isTesting {
// 		m := mockClockIn(mockClockInResp)
// 		return ctx.JSON(http.StatusOK, m)
// 	}

// 	result, err := TimeCardService.ClockInEmployee(employee)
// 	if err != nil {
// 		ctx.JSON(err.Status, err)
// 		return nil
// 	}

// 	return ctx.JSON(http.StatusOK, result)
// }

// func ClockOutHandler(ctx echo.Context, isTesting bool) error {
// 	employeeId, idErr := api.GetEmployeeId(ctx.Param("id"))
// 	if idErr != nil {
// 		ctx.JSON(idErr.Status, idErr)
// 		return nil
// 	}

// 	var employee database.TimeCard
// 	employee.EmployeeID = employeeId

// 	//improper use of unit test handling
// 	if isTesting {
// 		m := mockClockOut(mockClockOutResp)
// 		return ctx.JSON(http.StatusOK, m)
// 	}

// 	result, err := TimeCardService.ClockOutEmployee(employee)
// 	if err != nil {
// 		ctx.JSON(err.Status, err)
// 		return nil
// 	}

// 	return ctx.JSON(http.StatusOK, result)
// }

// func TotalTimeHandler(ctx echo.Context, isTesting bool) error {
// 	employeeId, idErr := api.GetEmployeeId(ctx.Param("id"))
// 	if idErr != nil {
// 		ctx.JSON(idErr.Status, idErr)
// 		return nil
// 	}

// 	//improper use of unit test handling
// 	if isTesting {
// 		m := mockTotalTime(mockTotalTimeResp)
// 		return ctx.JSON(http.StatusOK, m)
// 	}

// 	employee, getErr := TimeCardService.GetTotalTime(employeeId)
// 	if getErr != nil {
// 		ctx.JSON(getErr.Status, getErr)
// 		return nil
// 	}

// 	return ctx.JSON(http.StatusOK, employee)
// }
