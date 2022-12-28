package employee

import (
	"fmt"
	"log"
	"time"
	"timeCardSimple/database"
	"timeCardSimple/errors"
)

var (
	EmployeeService employeesServiceInterface = &employeesService{}
)

type employeesService struct{}

type employeesServiceInterface interface {
	CreateEmployee(database.Employee) (*database.Employee, *errors.RestErr)
	GetAllEmployees() (database.Employees, *errors.RestErr)
	GetEmployee(int64) (*database.Employee, *errors.RestErr)
	DeleteEmployee(int64) *errors.RestErr
	ClockInEmployee(database.Employee) (*database.Employee, *errors.RestErr)
	ClockOutEmployee(database.Employee) (*database.Employee, *errors.RestErr)
}

func (s *employeesService) CreateEmployee(employee database.Employee) (*database.Employee, *errors.RestErr) {

	if err := employee.Save(); err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s *employeesService) GetAllEmployees() (database.Employees, *errors.RestErr) {
	dao := &database.Employee{}
	return dao.GetAll()
}

func (s *employeesService) GetEmployee(employeeId int64) (*database.Employee, *errors.RestErr) {
	result := &database.Employee{EmployeeID: employeeId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *employeesService) DeleteEmployee(employeeId int64) *errors.RestErr {
	employee := &database.Employee{EmployeeID: employeeId}
	_, err := EmployeeService.GetEmployee(employee.EmployeeID)

	if err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("Employee ID %d does not exist", employeeId))
	}

	return employee.Delete()
}

func (s *employeesService) ClockInEmployee(employee database.Employee) (*database.Employee, *errors.RestErr) {
	current := &database.Employee{EmployeeID: employee.EmployeeID}
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

func (s *employeesService) ClockOutEmployee(employee database.Employee) (*database.Employee, *errors.RestErr) {
	current := &database.Employee{EmployeeID: employee.EmployeeID}
	if err := current.Get(); err != nil {
		return nil, err
	}

	current.ClockOut = time.Now().UTC().Format("Mon Jan _2 15:04:05 MST 2006")
	log.Printf("Employee clocked out at: %s", current.ClockOut)

	if err := current.EmployeeClockOut(); err != nil {
		return nil, err
	}

	// current.TotalTime = employeeTotalTime(current.ClockIn, current.ClockOut)
	// log.Printf("Employee: %s worked for a total of: %v", current.Name, current.TotalTime)

	return current, nil
}
