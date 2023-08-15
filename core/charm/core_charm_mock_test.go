// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/charm (interfaces: SelectorLogger,SelectorModelConfig)

// Package charm is a generated GoMock package.
package charm

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSelectorLogger is a mock of SelectorLogger interface.
type MockSelectorLogger struct {
	ctrl     *gomock.Controller
	recorder *MockSelectorLoggerMockRecorder
}

// MockSelectorLoggerMockRecorder is the mock recorder for MockSelectorLogger.
type MockSelectorLoggerMockRecorder struct {
	mock *MockSelectorLogger
}

// NewMockSelectorLogger creates a new mock instance.
func NewMockSelectorLogger(ctrl *gomock.Controller) *MockSelectorLogger {
	mock := &MockSelectorLogger{ctrl: ctrl}
	mock.recorder = &MockSelectorLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSelectorLogger) EXPECT() *MockSelectorLoggerMockRecorder {
	return m.recorder
}

// Infof mocks base method.
func (m *MockSelectorLogger) Infof(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockSelectorLoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockSelectorLogger)(nil).Infof), varargs...)
}

// Tracef mocks base method.
func (m *MockSelectorLogger) Tracef(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockSelectorLoggerMockRecorder) Tracef(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockSelectorLogger)(nil).Tracef), varargs...)
}

// MockSelectorModelConfig is a mock of SelectorModelConfig interface.
type MockSelectorModelConfig struct {
	ctrl     *gomock.Controller
	recorder *MockSelectorModelConfigMockRecorder
}

// MockSelectorModelConfigMockRecorder is the mock recorder for MockSelectorModelConfig.
type MockSelectorModelConfigMockRecorder struct {
	mock *MockSelectorModelConfig
}

// NewMockSelectorModelConfig creates a new mock instance.
func NewMockSelectorModelConfig(ctrl *gomock.Controller) *MockSelectorModelConfig {
	mock := &MockSelectorModelConfig{ctrl: ctrl}
	mock.recorder = &MockSelectorModelConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSelectorModelConfig) EXPECT() *MockSelectorModelConfigMockRecorder {
	return m.recorder
}

// DefaultBase mocks base method.
func (m *MockSelectorModelConfig) DefaultBase() (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultBase")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// DefaultBase indicates an expected call of DefaultBase.
func (mr *MockSelectorModelConfigMockRecorder) DefaultBase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultBase", reflect.TypeOf((*MockSelectorModelConfig)(nil).DefaultBase))
}
