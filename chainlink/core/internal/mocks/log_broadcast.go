// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	eth "github.com/smartcontractkit/chainlink/core/services/eth"
	mock "github.com/stretchr/testify/mock"
)

// LogBroadcast is an autogenerated mock type for the LogBroadcast type
type LogBroadcast struct {
	mock.Mock
}

// Log provides a mock function with given fields:
func (_m *LogBroadcast) Log() eth.Log {
	ret := _m.Called()

	var r0 eth.Log
	if rf, ok := ret.Get(0).(func() eth.Log); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(eth.Log)
		}
	}

	return r0
}

// MarkConsumed provides a mock function with given fields:
func (_m *LogBroadcast) MarkConsumed() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLog provides a mock function with given fields: _a0
func (_m *LogBroadcast) UpdateLog(_a0 eth.Log) {
	_m.Called(_a0)
}

// WasAlreadyConsumed provides a mock function with given fields:
func (_m *LogBroadcast) WasAlreadyConsumed() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
