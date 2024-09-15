package employees

import (
	"fmt"
	"timeCardSimple/api"
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
}

func (s *employeesService) CreateEmployee(employee database.Employee) (*database.Employee, *errors.RestErr) {

	employee.DateCreated = api.GetNowDBFormat()

	if err := database.Save(&employee); err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s *employeesService) GetAllEmployees() (database.Employees, *errors.RestErr) {
	return database.GetAll()
}

func (s *employeesService) GetEmployee(employeeId int64) (*database.Employee, *errors.RestErr) {
	result := &database.Employee{EmployeeID: employeeId}
	if err := database.Get(result); err != nil {
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

	return database.Delete(employee)
}
