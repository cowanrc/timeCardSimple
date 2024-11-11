package employeesvc

import (
	"context"
	"errors"
	"log"
	"testing"
	"timeCardSimple/App/domain/employee/employeetest"
	"timeCardSimple/app/domain/employee"

	"github.com/golang/mock/gomock"
)

var (
	firstName = "First"
	lastName  = "Last"
	email     = "test@email.com"
)

func createEmployeeParams() employee.CreateParams {
	return employee.CreateParams{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

func Test_CreateEmployee_AddRepoErr(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	errAddEmployee := errors.New("Add Employee Error")

	employeeRepo := employeetest.NewMockRepo(mc)
	employeeRepo.EXPECT().AddEmployee(gomock.Any(), gomock.Any()).Return(errAddEmployee)

	s := &Service{
		employeeRepo: employeeRepo,
	}

	createParams := createEmployeeParams()

	_, err := s.CreateEmployee(
		context.Background(),
		createParams,
	)

	if err != errAddEmployee {
		t.Errorf("incorrect error, received: %v, expected: %v", err, errAddEmployee)
	}
}

func Test_CreateEmployee_Success(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	employeeRepo := employeetest.NewMockRepo(mc)
	employeeRepo.EXPECT().AddEmployee(gomock.Any(), gomock.Any()).Return(nil)

	s := &Service{
		employeeRepo: employeeRepo,
	}

	createParams := createEmployeeParams()

	e, err := s.CreateEmployee(
		context.Background(),
		createParams,
	)

	if err != nil {
		t.Errorf("incorrect error, received: %v, expected: %v", err, nil)
	}

	if e == nil {
		t.Errorf("incorrect employee, recieved: %v, expected: %v", e, &employee.Employee{})
	}
}

func Test_GetEmployees_Sucess(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	createParams1 := createEmployeeParams()
	createParams2 := employee.CreateParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "fake@test.com",
	}

	employee1, err := employee.New(createParams1)
	if err != nil {
		log.Fatalf("Error creating employee1")
	}

	employee2, err := employee.New(createParams2)
	if err != nil {
		log.Fatalf("Error creating employee2")
	}

	expectedEmployees := []*employee.Employee{employee1, employee2}

	employeeRepo := employeetest.NewMockRepo(mc)
	employeeRepo.EXPECT().GetAllEmployees(gomock.Any()).Return(expectedEmployees, nil)

	s := &Service{
		employeeRepo: employeeRepo,
	}

	receivedEmployees, err := s.GetEmployees(
		context.Background(),
	)

	if err != nil {
		t.Errorf("incorrect error, received: %v, expected: %v", err, nil)
	}

	if len(receivedEmployees) != len(expectedEmployees) {
		t.Errorf("incorrect number of employees returned, receivced: %v, expected: %v", len(receivedEmployees), len(expectedEmployees))
	}
}
