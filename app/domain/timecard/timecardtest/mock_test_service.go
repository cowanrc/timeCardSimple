// Code generated by MockGen. DO NOT EDIT.
// Source: timeCardSimple/app/domain/timecard (interfaces: Service)

// Package timecardtest is a generated GoMock package.
package timecardtest

import (
	context "context"
	reflect "reflect"
	id "timeCardSimple/app/domain/id"
	timecard "timeCardSimple/app/domain/timecard"

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

// CreateEmployeeTimecard mocks base method.
func (m *MockService) CreateEmployeeTimecard(arg0 context.Context, arg1 id.ID) (*timecard.Timecard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployeeTimecard", arg0, arg1)
	ret0, _ := ret[0].(*timecard.Timecard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmployeeTimecard indicates an expected call of CreateEmployeeTimecard.
func (mr *MockServiceMockRecorder) CreateEmployeeTimecard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployeeTimecard", reflect.TypeOf((*MockService)(nil).CreateEmployeeTimecard), arg0, arg1)
}
