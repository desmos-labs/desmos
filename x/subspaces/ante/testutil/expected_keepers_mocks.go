// Code generated by MockGen. DO NOT EDIT.
// Source: ./x/subspaces/ante/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/desmos-labs/desmos/v4/x/subspaces/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(moduleName string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", moduleName)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), moduleName)
}

// GetParams mocks base method.
func (m *MockAccountKeeper) GetParams(ctx types.Context) types0.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types0.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockAccountKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockAccountKeeper)(nil).GetParams), ctx)
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(ctx types.Context, acc types0.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", ctx, acc)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(ctx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), ctx, acc)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// IsSendEnabledCoins mocks base method.
func (m *MockBankKeeper) IsSendEnabledCoins(ctx types.Context, coins ...types.Coin) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range coins {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsSendEnabledCoins", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsSendEnabledCoins indicates an expected call of IsSendEnabledCoins.
func (mr *MockBankKeeperMockRecorder) IsSendEnabledCoins(ctx interface{}, coins ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, coins...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSendEnabledCoins", reflect.TypeOf((*MockBankKeeper)(nil).IsSendEnabledCoins), varargs...)
}

// SendCoins mocks base method.
func (m *MockBankKeeper) SendCoins(ctx types.Context, from, to types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, from, to, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankKeeperMockRecorder) SendCoins(ctx, from, to, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankKeeper)(nil).SendCoins), ctx, from, to, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// MockSubspacesKeeper is a mock of SubspacesKeeper interface.
type MockSubspacesKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockSubspacesKeeperMockRecorder
}

// MockSubspacesKeeperMockRecorder is the mock recorder for MockSubspacesKeeper.
type MockSubspacesKeeperMockRecorder struct {
	mock *MockSubspacesKeeper
}

// NewMockSubspacesKeeper creates a new mock instance.
func NewMockSubspacesKeeper(ctrl *gomock.Controller) *MockSubspacesKeeper {
	mock := &MockSubspacesKeeper{ctrl: ctrl}
	mock.recorder = &MockSubspacesKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubspacesKeeper) EXPECT() *MockSubspacesKeeperMockRecorder {
	return m.recorder
}

// GetSubspace mocks base method.
func (m *MockSubspacesKeeper) GetSubspace(ctx types.Context, subspaceID uint64) (types1.Subspace, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubspace", ctx, subspaceID)
	ret0, _ := ret[0].(types1.Subspace)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetSubspace indicates an expected call of GetSubspace.
func (mr *MockSubspacesKeeperMockRecorder) GetSubspace(ctx, subspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubspace", reflect.TypeOf((*MockSubspacesKeeper)(nil).GetSubspace), ctx, subspaceID)
}

// UseGrantedFees mocks base method.
func (m *MockSubspacesKeeper) UseGrantedFees(ctx types.Context, subspaceID uint64, grantee types.AccAddress, fees types.Coins, msgs []types.Msg) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseGrantedFees", ctx, subspaceID, grantee, fees, msgs)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UseGrantedFees indicates an expected call of UseGrantedFees.
func (mr *MockSubspacesKeeperMockRecorder) UseGrantedFees(ctx, subspaceID, grantee, fees, msgs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseGrantedFees", reflect.TypeOf((*MockSubspacesKeeper)(nil).UseGrantedFees), ctx, subspaceID, grantee, fees, msgs)
}

// MockAuthDeductFeeDecorator is a mock of AuthDeductFeeDecorator interface.
type MockAuthDeductFeeDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthDeductFeeDecoratorMockRecorder
}

// MockAuthDeductFeeDecoratorMockRecorder is the mock recorder for MockAuthDeductFeeDecorator.
type MockAuthDeductFeeDecoratorMockRecorder struct {
	mock *MockAuthDeductFeeDecorator
}

// NewMockAuthDeductFeeDecorator creates a new mock instance.
func NewMockAuthDeductFeeDecorator(ctrl *gomock.Controller) *MockAuthDeductFeeDecorator {
	mock := &MockAuthDeductFeeDecorator{ctrl: ctrl}
	mock.recorder = &MockAuthDeductFeeDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthDeductFeeDecorator) EXPECT() *MockAuthDeductFeeDecoratorMockRecorder {
	return m.recorder
}

// AnteHandle mocks base method.
func (m *MockAuthDeductFeeDecorator) AnteHandle(ctx types.Context, tx types.Tx, simulate bool, next types.AnteHandler) (types.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnteHandle", ctx, tx, simulate, next)
	ret0, _ := ret[0].(types.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnteHandle indicates an expected call of AnteHandle.
func (mr *MockAuthDeductFeeDecoratorMockRecorder) AnteHandle(ctx, tx, simulate, next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnteHandle", reflect.TypeOf((*MockAuthDeductFeeDecorator)(nil).AnteHandle), ctx, tx, simulate, next)
}
