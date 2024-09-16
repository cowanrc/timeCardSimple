package api

import (
	"timeCardSimple/app/domain/employee"

	"github.com/gogolfing/httpmux"
)

type API struct {
	EmployeeSVC employee.Service
}

func (a *API) RegisterRoutes() *httpmux.Mux {
	mux := newMux()

	a.registerEmployee(mux.SubRoute("/employees"))

	return mux

}

func newMux() *httpmux.Mux {
	mux := httpmux.New()
	return mux
}
