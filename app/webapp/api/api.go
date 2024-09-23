package api

import (
	"timeCardSimple/app/domain/employee"

	"github.com/go-chi/chi"
)

type API struct {
	Repos       *Repos `valid:"required"`
	EmployeeSVC employee.Service
}

func (a *API) RegisterRoutes(r chi.Router) {
	r.Route("/employees", func(r chi.Router) {
		r.Post("/", a.CreateEmployee)

		r.Route("/{employeeID}", func(r chi.Router) {
			r.Use(a.EmployeeCTX)
			r.Get("/", a.GetEmployeeByID)
			// r.Delete("/", a.DeleteEmployee)
		})
	})

}
