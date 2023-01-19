// Code generated by MockGen. DO NOT EDIT.
// Source: chains/evm/calls/transactor/transact.go

// Package mock_transactor is a generated GoMock package.
package mock_transactor

import (
	reflect "reflect"

	transactor "github.com/nonceblox/chainbridge-core/chains/evm/calls/transactor"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
)

// MockTransactor is a mock of Transactor interface.
type MockTransactor struct {
	ctrl     *gomock.Controller
	recorder *MockTransactorMockRecorder
}

// MockTransactorMockRecorder is the mock recorder for MockTransactor.
type MockTransactorMockRecorder struct {
	mock *MockTransactor
}

// NewMockTransactor creates a new mock instance.
func NewMockTransactor(ctrl *gomock.Controller) *MockTransactor {
	mock := &MockTransactor{ctrl: ctrl}
	mock.recorder = &MockTransactorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactor) EXPECT() *MockTransactorMockRecorder {
	return m.recorder
}

// Transact mocks base method.
func (m *MockTransactor) Transact(to *common.Address, data []byte, opts transactor.TransactOptions) (*common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transact", to, data, opts)
	ret0, _ := ret[0].(*common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transact indicates an expected call of Transact.
func (mr *MockTransactorMockRecorder) Transact(to, data, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transact", reflect.TypeOf((*MockTransactor)(nil).Transact), to, data, opts)
}
