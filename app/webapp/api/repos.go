package api

import "timeCardSimple/app/domain/employee"

type Repos struct {
	Employees employee.Repo `valid:"required"`
}
