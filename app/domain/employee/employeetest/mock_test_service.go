// Code generated by MockGen. DO NOT EDIT.
// Source: timeCardSimple/app/domain/employee (interfaces: Service)

// Package employeetest is a generated GoMock package.
package employeetest

import (
	context "context"
	reflect "reflect"
	employee "timeCardSimple/app/domain/employee"
	id "timeCardSimple/app/domain/id"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateEmployee mocks base method.
func (m *MockService) CreateEmployee(arg0 context.Context, arg1 employee.CreateParams) (*employee.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployee", arg0, arg1)
	ret0, _ := ret[0].(*employee.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmployee indicates an expected call of CreateEmployee.
func (mr *MockServiceMockRecorder) CreateEmployee(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployee", reflect.TypeOf((*MockService)(nil).CreateEmployee), arg0, arg1)
}

// DeleteEmployee mocks base method.
func (m *MockService) DeleteEmployee(arg0 context.Context, arg1 id.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmployee", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmployee indicates an expected call of DeleteEmployee.
func (mr *MockServiceMockRecorder) DeleteEmployee(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmployee", reflect.TypeOf((*MockService)(nil).DeleteEmployee), arg0, arg1)
}

// GetEmployeeByID mocks base method.
func (m *MockService) GetEmployeeByID(arg0 context.Context, arg1 id.ID) (*employee.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeByID", arg0, arg1)
	ret0, _ := ret[0].(*employee.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeByID indicates an expected call of GetEmployeeByID.
func (mr *MockServiceMockRecorder) GetEmployeeByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeByID", reflect.TypeOf((*MockService)(nil).GetEmployeeByID), arg0, arg1)
}

// GetEmployees mocks base method.
func (m *MockService) GetEmployees(arg0 context.Context) (*[]employee.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployees", arg0)
	ret0, _ := ret[0].(*[]employee.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployees indicates an expected call of GetEmployees.
func (mr *MockServiceMockRecorder) GetEmployees(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployees", reflect.TypeOf((*MockService)(nil).GetEmployees), arg0)
}
