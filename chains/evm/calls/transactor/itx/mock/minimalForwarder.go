// Code generated by MockGen. DO NOT EDIT.
// Source: ./chains/evm/calls/transactor/itx/minimalForwarder.go

// Package mock_itx is a generated GoMock package.
package mock_itx

import (
	big "math/big"
	reflect "reflect"

	forwarder "github.com/nonceblox/chainbridge-core/chains/evm/calls/contracts/forwarder"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
)

// MockForwarderContract is a mock of ForwarderContract interface.
type MockForwarderContract struct {
	ctrl     *gomock.Controller
	recorder *MockForwarderContractMockRecorder
}

// MockForwarderContractMockRecorder is the mock recorder for MockForwarderContract.
type MockForwarderContractMockRecorder struct {
	mock *MockForwarderContract
}

// NewMockForwarderContract creates a new mock instance.
func NewMockForwarderContract(ctrl *gomock.Controller) *MockForwarderContract {
	mock := &MockForwarderContract{ctrl: ctrl}
	mock.recorder = &MockForwarderContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockForwarderContract) EXPECT() *MockForwarderContractMockRecorder {
	return m.recorder
}

// ContractAddress mocks base method.
func (m *MockForwarderContract) ContractAddress() *common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContractAddress")
	ret0, _ := ret[0].(*common.Address)
	return ret0
}

// ContractAddress indicates an expected call of ContractAddress.
func (mr *MockForwarderContractMockRecorder) ContractAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContractAddress", reflect.TypeOf((*MockForwarderContract)(nil).ContractAddress))
}

// GetNonce mocks base method.
func (m *MockForwarderContract) GetNonce(from common.Address) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", from)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNonce indicates an expected call of GetNonce.
func (mr *MockForwarderContractMockRecorder) GetNonce(from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockForwarderContract)(nil).GetNonce), from)
}

// PrepareExecute mocks base method.
func (m *MockForwarderContract) PrepareExecute(forwardReq forwarder.ForwardRequest, sig []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareExecute", forwardReq, sig)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareExecute indicates an expected call of PrepareExecute.
func (mr *MockForwarderContractMockRecorder) PrepareExecute(forwardReq, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareExecute", reflect.TypeOf((*MockForwarderContract)(nil).PrepareExecute), forwardReq, sig)
}

// MockNonceStorer is a mock of NonceStorer interface.
type MockNonceStorer struct {
	ctrl     *gomock.Controller
	recorder *MockNonceStorerMockRecorder
}

// MockNonceStorerMockRecorder is the mock recorder for MockNonceStorer.
type MockNonceStorerMockRecorder struct {
	mock *MockNonceStorer
}

// NewMockNonceStorer creates a new mock instance.
func NewMockNonceStorer(ctrl *gomock.Controller) *MockNonceStorer {
	mock := &MockNonceStorer{ctrl: ctrl}
	mock.recorder = &MockNonceStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNonceStorer) EXPECT() *MockNonceStorerMockRecorder {
	return m.recorder
}

// GetNonce mocks base method.
func (m *MockNonceStorer) GetNonce(chainID *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", chainID)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNonce indicates an expected call of GetNonce.
func (mr *MockNonceStorerMockRecorder) GetNonce(chainID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockNonceStorer)(nil).GetNonce), chainID)
}

// StoreNonce mocks base method.
func (m *MockNonceStorer) StoreNonce(chainID, nonce *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreNonce", chainID, nonce)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreNonce indicates an expected call of StoreNonce.
func (mr *MockNonceStorerMockRecorder) StoreNonce(chainID, nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreNonce", reflect.TypeOf((*MockNonceStorer)(nil).StoreNonce), chainID, nonce)
}
