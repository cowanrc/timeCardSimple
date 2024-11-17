package weeklysummary

import (
	"time"
	"timeCardSimple/app/domain/id"
)

func (ws *WeeklySummary) ID() id.ID {
	return ws.options.ID
}

func (ws *WeeklySummary) EmployeeID() id.ID {
	return ws.options.EmployeeID
}

func (ws *WeeklySummary) StartDate() *time.Time {
	return ws.options.StartDate
}

func (ws *WeeklySummary) DaysWorked() int64 {
	return ws.options.DaysWorked
}

func (ws *WeeklySummary) TotalHours() int64 {
	return ws.options.TotalHours
}
