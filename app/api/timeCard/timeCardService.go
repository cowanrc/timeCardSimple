package timeCard

// import (
// 	"time"
// 	"timeCardSimple/api"
// 	"timeCardSimple/database"
// 	"timeCardSimple/errors"
// )

// var (
// 	TimeCardService timeCardServiceInterface = &timeCardService{}
// )

// type timeCardService struct{}

// type timeCardServiceInterface interface {
// 	ClockInEmployee(database.TimeCard) (*database.TimeCard, *errors.RestErr)
// 	ClockOutEmployee(database.TimeCard) (*database.TimeCard, *errors.RestErr)
// 	GetTotalTime(int64) (*database.TimeCard, *errors.RestErr)
// }

// func (s *timeCardService) ClockInEmployee(employee database.TimeCard) (*database.TimeCard, *errors.RestErr) {
// 	current := &database.TimeCard{EmployeeID: employee.EmployeeID}
// 	if err := database.GetSimple(current); err != nil {
// 		return nil, err
// 	}

// 	current.ClockIn = time.Now().UTC().Format("Mon Jan _2 15:04:05 MST 2006")

// 	if err := database.EmployeeClockIn(current); err != nil {
// 		return nil, err
// 	}

// 	return current, nil
// }

// func (s *timeCardService) ClockOutEmployee(employee database.TimeCard) (*database.TimeCard, *errors.RestErr) {
// 	current := &database.TimeCard{EmployeeID: employee.EmployeeID}
// 	if err := database.GetClockIn(current); err != nil {
// 		return nil, err
// 	}

// 	current.ClockOut = time.Now().UTC().Format("Mon Jan _2 15:04:05 MST 2006")

// 	if err := database.EmployeeClockOut(current); err != nil {
// 		return nil, err
// 	}

// 	return current, nil
// }

// func (s *timeCardService) GetTotalTime(employeeId int64) (*database.TimeCard, *errors.RestErr) {
// 	employee := &database.TimeCard{EmployeeID: employeeId}
// 	if err := database.GetClockInClockOut(employee); err != nil {
// 		return nil, err
// 	}

// 	employee.TotalTime = api.CalculateTotalTime(employee.ClockIn, employee.ClockOut)

// 	if err := database.EmployeeTotalTime(employee); err != nil {
// 		return nil, err
// 	}

// 	if err := database.GetTime(employee); err != nil {
// 		return nil, err
// 	}

// 	return employee, nil
// }
