// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/surahman/FTeX/pkg/postgres (interfaces: Postgres)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	models "github.com/surahman/FTeX/pkg/models/postgres"
)

// MockPostgres is a mock of Postgres interface.
type MockPostgres struct {
	ctrl     *gomock.Controller
	recorder *MockPostgresMockRecorder
}

// MockPostgresMockRecorder is the mock recorder for MockPostgres.
type MockPostgresMockRecorder struct {
	mock *MockPostgres
}

// NewMockPostgres creates a new mock instance.
func NewMockPostgres(ctrl *gomock.Controller) *MockPostgres {
	mock := &MockPostgres{ctrl: ctrl}
	mock.recorder = &MockPostgresMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostgres) EXPECT() *MockPostgresMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPostgres) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPostgresMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPostgres)(nil).Close))
}

// Healthcheck mocks base method.
func (m *MockPostgres) Healthcheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Healthcheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// Healthcheck indicates an expected call of Healthcheck.
func (mr *MockPostgresMockRecorder) Healthcheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthcheck", reflect.TypeOf((*MockPostgres)(nil).Healthcheck))
}

// Open mocks base method.
func (m *MockPostgres) Open() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *MockPostgresMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockPostgres)(nil).Open))
}

// UserCredentials mocks base method.
func (m *MockPostgres) UserCredentials(arg0 string) (uuid.UUID, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserCredentials", arg0)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserCredentials indicates an expected call of UserCredentials.
func (mr *MockPostgresMockRecorder) UserCredentials(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCredentials", reflect.TypeOf((*MockPostgres)(nil).UserCredentials), arg0)
}

// UserGetInfo mocks base method.
func (m *MockPostgres) UserGetInfo(arg0 uuid.UUID) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserGetInfo", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserGetInfo indicates an expected call of UserGetInfo.
func (mr *MockPostgresMockRecorder) UserGetInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGetInfo", reflect.TypeOf((*MockPostgres)(nil).UserGetInfo), arg0)
}

// UserRegister mocks base method.
func (m *MockPostgres) UserRegister(arg0 *models.UserAccount) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserRegister", arg0)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserRegister indicates an expected call of UserRegister.
func (mr *MockPostgresMockRecorder) UserRegister(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRegister", reflect.TypeOf((*MockPostgres)(nil).UserRegister), arg0)
}
