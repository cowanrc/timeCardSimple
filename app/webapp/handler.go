package webapp

import (
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/employee/employeesvc"
	"timeCardSimple/app/domain/timecard"
	"timeCardSimple/app/domain/timecard/timecardsvc"
	"timeCardSimple/app/webapp/api"
)

func BuildRoot(
	repos *Repos,
) (*api.API, error) {
	employeeService := employeesvc.New(
		repos.Employee,
		repos.Timecard,
	)

	timecardService := timecardsvc.New(
		repos.Timecard,
		repos.Employee,
		repos.WeeklySummary,
		repos.PayPeriod,
	)

	api, err := buildAPI(
		repos,
		employeeService,
		timecardService,
	)

	if err != nil {
		return nil, err
	}

	return api, nil
}

func buildAPI(
	repos *Repos,
	employeeService employee.Service,
	timecardService timecard.Service,
) (*api.API, error) {

	a := &api.API{
		Repos: &api.Repos{
			Employees:     repos.Employee,
			Timecard:      repos.Timecard,
			WeeklySummary: repos.WeeklySummary,
			PayPeriod:     repos.PayPeriod,
		},
		EmployeeSVC: employeeService,
		TimecardSVC: timecardService,
	}

	return a, nil
}
