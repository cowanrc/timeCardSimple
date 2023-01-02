package timeCard

import (
	"log"
	"time"
	"timeCardSimple/api"
	"timeCardSimple/database"
	"timeCardSimple/errors"
)

var (
	TimeCardService timeCardServiceInterface = &timeCardService{}
)

type timeCardService struct{}

type timeCardServiceInterface interface {
	ClockInEmployee(database.TimeCard) (*database.TimeCard, *errors.RestErr)
	ClockOutEmployee(database.TimeCard) (*database.TimeCard, *errors.RestErr)
	GetTotalTime(int64) (*database.TimeCard, *errors.RestErr)
}

func (s *timeCardService) ClockInEmployee(employee database.TimeCard) (*database.TimeCard, *errors.RestErr) {
	current := &database.TimeCard{EmployeeID: employee.EmployeeID}
	if err := current.Get(); err != nil {
		return nil, err
	}

	current.ClockIn = time.Now().UTC().Format("Mon Jan _2 15:04:05 MST 2006")
	log.Printf("Employee clocked in at: %s", current.ClockIn)

	if err := current.EmployeeClockIn(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *timeCardService) ClockOutEmployee(employee database.TimeCard) (*database.TimeCard, *errors.RestErr) {
	current := &database.TimeCard{EmployeeID: employee.EmployeeID}
	if err := current.GetClockIn(); err != nil {
		return nil, err
	}

	current.ClockOut = time.Now().UTC().Format("Mon Jan _2 15:04:05 MST 2006")

	if err := current.EmployeeClockOut(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *timeCardService) GetTotalTime(employeeId int64) (*database.TimeCard, *errors.RestErr) {
	employee := &database.TimeCard{EmployeeID: employeeId}
	if err := employee.GetClockInClockOut(); err != nil {
		return nil, err
	}

	employee.TotalTime = api.CalculateTotalTime(employee.ClockIn, employee.ClockOut)

	if err := employee.EmployeeTotalTime(); err != nil {
		return nil, err
	}

	if err := employee.GetTime(); err != nil {
		return nil, err
	}

	return employee, nil
}
