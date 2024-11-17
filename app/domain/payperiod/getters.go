package payperiod

import (
	"time"
	"timeCardSimple/app/domain/id"
)

func (pp *PayPeriod) ID() id.ID {
	return pp.options.ID
}

func (pp *PayPeriod) EmployeeID() id.ID {
	return pp.options.EmployeeID
}

func (pp *PayPeriod) StartDate() *time.Time {
	return pp.options.StartDate
}

func (pp *PayPeriod) EndDate() *time.Time {
	return pp.options.EndDate
}

func (pp *PayPeriod) DaysWorked() uint64 {
	return pp.options.DaysWorked
}

func (pp *PayPeriod) TotalHours() uint64 {
	return pp.options.TotalHours
}
