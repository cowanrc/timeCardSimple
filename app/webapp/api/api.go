package api

import (
	"context"
	"net/http"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/timecard"

	"github.com/go-chi/chi"
)

type employeeContextKey struct{}

type API struct {
	Repos       *Repos `valid:"required"`
	EmployeeSVC employee.Service
	TimecardSVC timecard.Service
}

func (a *API) RegisterRoutes(r chi.Router) {
	r.Route("/employees", func(r chi.Router) {
		r.Post("/", a.CreateEmployee)
		r.Get("/", a.GetEmployees)

		r.Route("/{employeeID}", func(r chi.Router) {
			r.Use(a.EmployeeCTX)
			r.Get("/", a.GetEmployeeByID)
			r.Delete("/", a.DeleteEmployee)

			r.Route("/timecard", func(r chi.Router) {
				r.Post("/clock-in", a.ClockIn)
				// r.Post("/clock-out", a.ClockOut)
			})
		})
	})

}

func (a *API) EmployeeCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employeeID, err := id.ParseString(chi.URLParam(r, "employeeID"))
		if err != nil {
			NewBadRequestError("Error parsing employeeID")
			return
		}

		employee, err := a.EmployeeSVC.GetEmployeeByID(r.Context(), employeeID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), employeeContextKey{}, employee)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
