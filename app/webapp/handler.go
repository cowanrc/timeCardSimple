package webapp

import (
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/employee/employeesvc"
	"timeCardSimple/app/webapp/api"
)

func BuildRoot(
	repos *Repos,
) (*api.API, error) {
	employeeService := employeesvc.New(repos.Employee)

	api, err := buildAPI(
		repos,
		employeeService,
	)

	if err != nil {
		return nil, err
	}

	return api, nil
}

func buildAPI(
	repos *Repos,
	employeeService employee.Service,
) (*api.API, error) {

	a := &api.API{
		Repos: &api.Repos{
			Employees: repos.Employee,
		},
		EmployeeSVC: employeeService,
	}

	return a, nil
}
