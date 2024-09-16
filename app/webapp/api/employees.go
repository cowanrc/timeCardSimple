package api

import (
	"net/http"

	"github.com/gogolfing/httpmux"
)

const employeeID = "employee_id"

func (a *API) registerEmployee(route *httpmux.Route) *httpmux.Route {
	route.
		PostFunc(a.createEmployee)

	employeeRoute := route.SubRoute("/:" + employeeID).
		DeleteFunc(a.deleteEmployee)

	return employeeRoute

}

func (a *API) createEmployee(w http.ResponseWriter, r *http.Request) {

}

func (a *API) deleteEmployee(w http.ResponseWriter, r *http.Request) {

}
