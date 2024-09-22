package api

import (
	"net/http"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/webapp/api/dto"
	"timeCardSimple/app/webapp/api/form"

	"github.com/labstack/echo/v4"
)

func (a *API) CreateEmployeeHandler(ctx echo.Context) error {
	f := &form.CreateEmployee{}

	if err := ctx.Bind(f); err != nil {
		restErr := NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
	}

	newE := employee.CreateParams{
		FirstName: f.FirstName,
		LastName:  f.LastName,
		Email:     f.Email,
	}

	newEmployee, err := a.EmployeeSVC.CreateEmployee(ctx.Request().Context(), newE)
	if err != nil {
		restErr := NewBadRequestError("error creating a employee")
		ctx.JSON(restErr.Status, restErr)
	}

	transfer, err := dto.Transform(newEmployee)
	if err != nil {
		restErr := NewBadRequestError("error creating a employee")
		ctx.JSON(restErr.Status, restErr)
	}

	return ctx.JSON(http.StatusCreated, transfer)
}
