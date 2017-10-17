// Code generated by MockGen. DO NOT EDIT.
// Source: mockObject.go

// Package main is a generated GoMock package.
package mockExample

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDBService is a mock of DBService interface
type MockDBService struct {
	ctrl     *gomock.Controller
	recorder *MockDBServiceMockRecorder
}

// MockDBServiceMockRecorder is the mock recorder for MockDBService
type MockDBServiceMockRecorder struct {
	mock *MockDBService
}

// NewMockDBService creates a new mock instance
func NewMockDBService(ctrl *gomock.Controller) *MockDBService {
	mock := &MockDBService{ctrl: ctrl}
	mock.recorder = &MockDBServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDBService) EXPECT() *MockDBServiceMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockDBService) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockDBServiceMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockDBService)(nil).Init))
}

// Insert mocks base method
func (m *MockDBService) Insert(arg0 string) error {
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockDBServiceMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDBService)(nil).Insert), arg0)
}

// Get mocks base method
func (m *MockDBService) Get(arg0 string) error {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockDBServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDBService)(nil).Get), arg0)
}

// Count mocks base method
func (m *MockDBService) Count() error {
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(error)
	return ret0
}

// Count indicates an expected call of Count
func (mr *MockDBServiceMockRecorder) Count() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDBService)(nil).Count))
}

// DeleteAll mocks base method
func (m *MockDBService) DeleteAll() error {
	ret := m.ctrl.Call(m, "DeleteAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll
func (mr *MockDBServiceMockRecorder) DeleteAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockDBService)(nil).DeleteAll))
}
