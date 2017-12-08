// Code generated by MockGen. DO NOT EDIT.
// Source: ../models.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConfigStore is a mock of ConfigStore interface
type MockConfigStore struct {
	ctrl     *gomock.Controller
	recorder *MockConfigStoreMockRecorder
}

// MockConfigStoreMockRecorder is the mock recorder for MockConfigStore
type MockConfigStoreMockRecorder struct {
	mock *MockConfigStore
}

// NewMockConfigStore creates a new mock instance
func NewMockConfigStore(ctrl *gomock.Controller) *MockConfigStore {
	mock := &MockConfigStore{ctrl: ctrl}
	mock.recorder = &MockConfigStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigStore) EXPECT() *MockConfigStoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockConfigStore) Get(Type, ID string, result interface{}) error {
	ret := m.ctrl.Call(m, "Get", Type, ID, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockConfigStoreMockRecorder) Get(Type, ID, result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfigStore)(nil).Get), Type, ID, result)
}

// Set mocks base method
func (m *MockConfigStore) Set(Type, ID string, model interface{}) error {
	ret := m.ctrl.Call(m, "Set", Type, ID, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockConfigStoreMockRecorder) Set(Type, ID, model interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockConfigStore)(nil).Set), Type, ID, model)
}