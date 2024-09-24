package api

import (
	"context"
	"encoding/json"
	"net/http"

	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/webapp/api/dto"
	"timeCardSimple/app/webapp/api/form"

	"github.com/go-chi/chi"
)

type employeeContextKey struct{}

// EmployeCTX gets a specific employee with ID to be used in other subroutes
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

func (a *API) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	f := &form.CreateEmployee{}

	if err := json.NewDecoder(r.Body).Decode(f); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	newE := employee.CreateParams{
		FirstName: f.FirstName,
		LastName:  f.LastName,
		Email:     f.Email,
	}

	newEmployee, err := a.EmployeeSVC.CreateEmployee(ctx, newE)
	if err != nil {
		http.Error(w, "error creating employee", http.StatusBadRequest)
		return
	}

	transfer, err := dto.Transform(newEmployee)
	if err != nil {
		http.Error(w, "error transforming employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(transfer); err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}

func (a *API) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	employee, ok := ctx.Value(employeeContextKey{}).(*employee.Employee)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	transfer, err := dto.Transform(employee)
	if err != nil {
		http.Error(w, "error transforming employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transfer); err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}

func (a *API) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	employeeID, err := id.ParseString(chi.URLParam(r, "employeeID"))
	if err != nil {
		NewBadRequestError("Error parsing employeeID")
		return
	}

	if err = a.EmployeeSVC.DeleteEmployee(ctx, employeeID); err != nil {
		http.Error(w, "error deleting employee", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusAccepted)
}
