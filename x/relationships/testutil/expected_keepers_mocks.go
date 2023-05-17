// Code generated by MockGen. DO NOT EDIT.
// Source: ./x/relationships/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	gomock "github.com/golang/mock/gomock"
)

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

// GetAllSubspaces mocks base method.
func (m *MockSubspacesKeeper) GetAllSubspaces(ctx types.Context) []types0.Subspace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubspaces", ctx)
	ret0, _ := ret[0].([]types0.Subspace)
	return ret0
}

// GetAllSubspaces indicates an expected call of GetAllSubspaces.
func (mr *MockSubspacesKeeperMockRecorder) GetAllSubspaces(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubspaces", reflect.TypeOf((*MockSubspacesKeeper)(nil).GetAllSubspaces), ctx)
}

// HasSubspace mocks base method.
func (m *MockSubspacesKeeper) HasSubspace(ctx types.Context, subspaceID uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSubspace", ctx, subspaceID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSubspace indicates an expected call of HasSubspace.
func (mr *MockSubspacesKeeperMockRecorder) HasSubspace(ctx, subspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSubspace", reflect.TypeOf((*MockSubspacesKeeper)(nil).HasSubspace), ctx, subspaceID)
}
