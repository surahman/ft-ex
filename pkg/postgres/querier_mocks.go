// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/surahman/FTeX/pkg/postgres (interfaces: Querier)

// Package postgres is a generated GoMock package.
package postgres

import (
	context "context"
	reflect "reflect"

	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// cryptoCreateAccount mocks base method.
func (m *MockQuerier) cryptoCreateAccount(arg0 context.Context, arg1 *cryptoCreateAccountParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cryptoCreateAccount", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// cryptoCreateAccount indicates an expected call of cryptoCreateAccount.
func (mr *MockQuerierMockRecorder) cryptoCreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cryptoCreateAccount", reflect.TypeOf((*MockQuerier)(nil).cryptoCreateAccount), arg0, arg1)
}

// cryptoGetAccount mocks base method.
func (m *MockQuerier) cryptoGetAccount(arg0 context.Context, arg1 *cryptoGetAccountParams) (CryptoAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cryptoGetAccount", arg0, arg1)
	ret0, _ := ret[0].(CryptoAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// cryptoGetAccount indicates an expected call of cryptoGetAccount.
func (mr *MockQuerierMockRecorder) cryptoGetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cryptoGetAccount", reflect.TypeOf((*MockQuerier)(nil).cryptoGetAccount), arg0, arg1)
}

// cryptoGetAllAccounts mocks base method.
func (m *MockQuerier) cryptoGetAllAccounts(arg0 context.Context, arg1 *cryptoGetAllAccountsParams) ([]CryptoAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cryptoGetAllAccounts", arg0, arg1)
	ret0, _ := ret[0].([]CryptoAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// cryptoGetAllAccounts indicates an expected call of cryptoGetAllAccounts.
func (mr *MockQuerierMockRecorder) cryptoGetAllAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cryptoGetAllAccounts", reflect.TypeOf((*MockQuerier)(nil).cryptoGetAllAccounts), arg0, arg1)
}

// cryptoGetAllJournalTransactionPaginated mocks base method.
func (m *MockQuerier) cryptoGetAllJournalTransactionPaginated(arg0 context.Context, arg1 *cryptoGetAllJournalTransactionPaginatedParams) ([]CryptoJournal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cryptoGetAllJournalTransactionPaginated", arg0, arg1)
	ret0, _ := ret[0].([]CryptoJournal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// cryptoGetAllJournalTransactionPaginated indicates an expected call of cryptoGetAllJournalTransactionPaginated.
func (mr *MockQuerierMockRecorder) cryptoGetAllJournalTransactionPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cryptoGetAllJournalTransactionPaginated", reflect.TypeOf((*MockQuerier)(nil).cryptoGetAllJournalTransactionPaginated), arg0, arg1)
}

// cryptoGetJournalTransaction mocks base method.
func (m *MockQuerier) cryptoGetJournalTransaction(arg0 context.Context, arg1 *cryptoGetJournalTransactionParams) ([]CryptoJournal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cryptoGetJournalTransaction", arg0, arg1)
	ret0, _ := ret[0].([]CryptoJournal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// cryptoGetJournalTransaction indicates an expected call of cryptoGetJournalTransaction.
func (mr *MockQuerierMockRecorder) cryptoGetJournalTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cryptoGetJournalTransaction", reflect.TypeOf((*MockQuerier)(nil).cryptoGetJournalTransaction), arg0, arg1)
}

// cryptoPurchase mocks base method.
func (m *MockQuerier) cryptoPurchase(arg0 context.Context, arg1 *cryptoPurchaseParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cryptoPurchase", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// cryptoPurchase indicates an expected call of cryptoPurchase.
func (mr *MockQuerierMockRecorder) cryptoPurchase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cryptoPurchase", reflect.TypeOf((*MockQuerier)(nil).cryptoPurchase), arg0, arg1)
}

// cryptoSell mocks base method.
func (m *MockQuerier) cryptoSell(arg0 context.Context, arg1 *cryptoSellParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cryptoSell", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// cryptoSell indicates an expected call of cryptoSell.
func (mr *MockQuerierMockRecorder) cryptoSell(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cryptoSell", reflect.TypeOf((*MockQuerier)(nil).cryptoSell), arg0, arg1)
}

// fiatCreateAccount mocks base method.
func (m *MockQuerier) fiatCreateAccount(arg0 context.Context, arg1 *fiatCreateAccountParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fiatCreateAccount", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fiatCreateAccount indicates an expected call of fiatCreateAccount.
func (mr *MockQuerierMockRecorder) fiatCreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fiatCreateAccount", reflect.TypeOf((*MockQuerier)(nil).fiatCreateAccount), arg0, arg1)
}

// fiatExternalTransferJournalEntry mocks base method.
func (m *MockQuerier) fiatExternalTransferJournalEntry(arg0 context.Context, arg1 *fiatExternalTransferJournalEntryParams) (fiatExternalTransferJournalEntryRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fiatExternalTransferJournalEntry", arg0, arg1)
	ret0, _ := ret[0].(fiatExternalTransferJournalEntryRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fiatExternalTransferJournalEntry indicates an expected call of fiatExternalTransferJournalEntry.
func (mr *MockQuerierMockRecorder) fiatExternalTransferJournalEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fiatExternalTransferJournalEntry", reflect.TypeOf((*MockQuerier)(nil).fiatExternalTransferJournalEntry), arg0, arg1)
}

// fiatGetAccount mocks base method.
func (m *MockQuerier) fiatGetAccount(arg0 context.Context, arg1 *fiatGetAccountParams) (FiatAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fiatGetAccount", arg0, arg1)
	ret0, _ := ret[0].(FiatAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fiatGetAccount indicates an expected call of fiatGetAccount.
func (mr *MockQuerierMockRecorder) fiatGetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fiatGetAccount", reflect.TypeOf((*MockQuerier)(nil).fiatGetAccount), arg0, arg1)
}

// fiatGetAllAccounts mocks base method.
func (m *MockQuerier) fiatGetAllAccounts(arg0 context.Context, arg1 *fiatGetAllAccountsParams) ([]FiatAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fiatGetAllAccounts", arg0, arg1)
	ret0, _ := ret[0].([]FiatAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fiatGetAllAccounts indicates an expected call of fiatGetAllAccounts.
func (mr *MockQuerierMockRecorder) fiatGetAllAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fiatGetAllAccounts", reflect.TypeOf((*MockQuerier)(nil).fiatGetAllAccounts), arg0, arg1)
}

// fiatGetAllJournalTransactionPaginated mocks base method.
func (m *MockQuerier) fiatGetAllJournalTransactionPaginated(arg0 context.Context, arg1 *fiatGetAllJournalTransactionPaginatedParams) ([]FiatJournal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fiatGetAllJournalTransactionPaginated", arg0, arg1)
	ret0, _ := ret[0].([]FiatJournal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fiatGetAllJournalTransactionPaginated indicates an expected call of fiatGetAllJournalTransactionPaginated.
func (mr *MockQuerierMockRecorder) fiatGetAllJournalTransactionPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fiatGetAllJournalTransactionPaginated", reflect.TypeOf((*MockQuerier)(nil).fiatGetAllJournalTransactionPaginated), arg0, arg1)
}

// fiatGetJournalTransaction mocks base method.
func (m *MockQuerier) fiatGetJournalTransaction(arg0 context.Context, arg1 *fiatGetJournalTransactionParams) ([]FiatJournal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fiatGetJournalTransaction", arg0, arg1)
	ret0, _ := ret[0].([]FiatJournal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fiatGetJournalTransaction indicates an expected call of fiatGetJournalTransaction.
func (mr *MockQuerierMockRecorder) fiatGetJournalTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fiatGetJournalTransaction", reflect.TypeOf((*MockQuerier)(nil).fiatGetJournalTransaction), arg0, arg1)
}

// fiatGetJournalTransactionForAccount mocks base method.
func (m *MockQuerier) fiatGetJournalTransactionForAccount(arg0 context.Context, arg1 *fiatGetJournalTransactionForAccountParams) ([]FiatJournal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fiatGetJournalTransactionForAccount", arg0, arg1)
	ret0, _ := ret[0].([]FiatJournal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fiatGetJournalTransactionForAccount indicates an expected call of fiatGetJournalTransactionForAccount.
func (mr *MockQuerierMockRecorder) fiatGetJournalTransactionForAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fiatGetJournalTransactionForAccount", reflect.TypeOf((*MockQuerier)(nil).fiatGetJournalTransactionForAccount), arg0, arg1)
}

// fiatInternalTransferJournalEntry mocks base method.
func (m *MockQuerier) fiatInternalTransferJournalEntry(arg0 context.Context, arg1 *fiatInternalTransferJournalEntryParams) (fiatInternalTransferJournalEntryRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fiatInternalTransferJournalEntry", arg0, arg1)
	ret0, _ := ret[0].(fiatInternalTransferJournalEntryRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fiatInternalTransferJournalEntry indicates an expected call of fiatInternalTransferJournalEntry.
func (mr *MockQuerierMockRecorder) fiatInternalTransferJournalEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fiatInternalTransferJournalEntry", reflect.TypeOf((*MockQuerier)(nil).fiatInternalTransferJournalEntry), arg0, arg1)
}

// fiatRowLockAccount mocks base method.
func (m *MockQuerier) fiatRowLockAccount(arg0 context.Context, arg1 *fiatRowLockAccountParams) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fiatRowLockAccount", arg0, arg1)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fiatRowLockAccount indicates an expected call of fiatRowLockAccount.
func (mr *MockQuerierMockRecorder) fiatRowLockAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fiatRowLockAccount", reflect.TypeOf((*MockQuerier)(nil).fiatRowLockAccount), arg0, arg1)
}

// fiatUpdateAccountBalance mocks base method.
func (m *MockQuerier) fiatUpdateAccountBalance(arg0 context.Context, arg1 *fiatUpdateAccountBalanceParams) (fiatUpdateAccountBalanceRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fiatUpdateAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(fiatUpdateAccountBalanceRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fiatUpdateAccountBalance indicates an expected call of fiatUpdateAccountBalance.
func (mr *MockQuerierMockRecorder) fiatUpdateAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fiatUpdateAccountBalance", reflect.TypeOf((*MockQuerier)(nil).fiatUpdateAccountBalance), arg0, arg1)
}

// testRoundHalfEven mocks base method.
func (m *MockQuerier) testRoundHalfEven(arg0 context.Context, arg1 *testRoundHalfEvenParams) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "testRoundHalfEven", arg0, arg1)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// testRoundHalfEven indicates an expected call of testRoundHalfEven.
func (mr *MockQuerierMockRecorder) testRoundHalfEven(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "testRoundHalfEven", reflect.TypeOf((*MockQuerier)(nil).testRoundHalfEven), arg0, arg1)
}

// userCreate mocks base method.
func (m *MockQuerier) userCreate(arg0 context.Context, arg1 *userCreateParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "userCreate", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// userCreate indicates an expected call of userCreate.
func (mr *MockQuerierMockRecorder) userCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "userCreate", reflect.TypeOf((*MockQuerier)(nil).userCreate), arg0, arg1)
}

// userDelete mocks base method.
func (m *MockQuerier) userDelete(arg0 context.Context, arg1 uuid.UUID) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "userDelete", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// userDelete indicates an expected call of userDelete.
func (mr *MockQuerierMockRecorder) userDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "userDelete", reflect.TypeOf((*MockQuerier)(nil).userDelete), arg0, arg1)
}

// userGetClientId mocks base method.
func (m *MockQuerier) userGetClientId(arg0 context.Context, arg1 string) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "userGetClientId", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// userGetClientId indicates an expected call of userGetClientId.
func (mr *MockQuerierMockRecorder) userGetClientId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "userGetClientId", reflect.TypeOf((*MockQuerier)(nil).userGetClientId), arg0, arg1)
}

// userGetCredentials mocks base method.
func (m *MockQuerier) userGetCredentials(arg0 context.Context, arg1 string) (userGetCredentialsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "userGetCredentials", arg0, arg1)
	ret0, _ := ret[0].(userGetCredentialsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// userGetCredentials indicates an expected call of userGetCredentials.
func (mr *MockQuerierMockRecorder) userGetCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "userGetCredentials", reflect.TypeOf((*MockQuerier)(nil).userGetCredentials), arg0, arg1)
}

// userGetInfo mocks base method.
func (m *MockQuerier) userGetInfo(arg0 context.Context, arg1 uuid.UUID) (userGetInfoRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "userGetInfo", arg0, arg1)
	ret0, _ := ret[0].(userGetInfoRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// userGetInfo indicates an expected call of userGetInfo.
func (mr *MockQuerierMockRecorder) userGetInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "userGetInfo", reflect.TypeOf((*MockQuerier)(nil).userGetInfo), arg0, arg1)
}
