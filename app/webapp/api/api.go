package api

import (
	"timeCardSimple/app/domain/employee"

	"github.com/labstack/echo/v4"
)

type API struct {
	Repos       *Repos `valid:"required"`
	EmployeeSVC employee.Service
}

func (a *API) RegisterRoutes(e *echo.Echo) {
	e.POST("/employees", a.CreateEmployeeHandler)

}
