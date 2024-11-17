package api

import (
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/payperiod"
	"timeCardSimple/app/domain/timecard"
	"timeCardSimple/app/domain/weeklysummary"
)

type Repos struct {
	Employees     employee.Repo      `valid:"required"`
	Timecard      timecard.Repo      `valid:"required"`
	WeeklySummary weeklysummary.Repo `valid:"required"`
	PayPeriod     payperiod.Repo     `valid:"required"`
}
