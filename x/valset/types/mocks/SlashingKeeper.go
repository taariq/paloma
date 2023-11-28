// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// SlashingKeeper is an autogenerated mock type for the SlashingKeeper type
type SlashingKeeper struct {
	mock.Mock
}

// Jail provides a mock function with given fields: _a0, _a1
func (_m *SlashingKeeper) Jail(_a0 types.Context, _a1 types.ConsAddress) {
	_m.Called(_a0, _a1)
}

// JailUntil provides a mock function with given fields: _a0, _a1, _a2
func (_m *SlashingKeeper) JailUntil(_a0 types.Context, _a1 types.ConsAddress, _a2 time.Time) {
	_m.Called(_a0, _a1, _a2)
}

// NewSlashingKeeper creates a new instance of SlashingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSlashingKeeper(t interface {
	mock.TestingT
	Cleanup(func())
},
) *SlashingKeeper {
	mock := &SlashingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
