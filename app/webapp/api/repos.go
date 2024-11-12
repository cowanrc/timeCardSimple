package api

import (
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/timecard"
)

type Repos struct {
	Employees employee.Repo `valid:"required"`
	Timecard  timecard.Repo `valid:"required"`
}
