// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/docker/registry/internal (interfaces: Initializer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	internal "github.com/juju/juju/docker/registry/internal"
)

// MockInitializer is a mock of Initializer interface.
type MockInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockInitializerMockRecorder
}

// MockInitializerMockRecorder is the mock recorder for MockInitializer.
type MockInitializerMockRecorder struct {
	mock *MockInitializer
}

// NewMockInitializer creates a new mock instance.
func NewMockInitializer(ctrl *gomock.Controller) *MockInitializer {
	mock := &MockInitializer{ctrl: ctrl}
	mock.recorder = &MockInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInitializer) EXPECT() *MockInitializerMockRecorder {
	return m.recorder
}

// DecideBaseURL mocks base method.
func (m *MockInitializer) DecideBaseURL() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecideBaseURL")
	ret0, _ := ret[0].(error)
	return ret0
}

// DecideBaseURL indicates an expected call of DecideBaseURL.
func (mr *MockInitializerMockRecorder) DecideBaseURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecideBaseURL", reflect.TypeOf((*MockInitializer)(nil).DecideBaseURL))
}

// WrapTransport mocks base method.
func (m *MockInitializer) WrapTransport(arg0 ...internal.TransportWrapper) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WrapTransport", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WrapTransport indicates an expected call of WrapTransport.
func (mr *MockInitializerMockRecorder) WrapTransport(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WrapTransport", reflect.TypeOf((*MockInitializer)(nil).WrapTransport), arg0...)
}
