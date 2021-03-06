// Automatically generated by MockGen. DO NOT EDIT!
// Source: laboratory.go

package laboratory

import (
	gomock "github.com/golang/mock/gomock"
)

// MockCreatures - Mock of Creatures interface
type MockCreatures struct {
	ctrl     *gomock.Controller
	recorder *_MockCreaturesRecorder
}

// Recorder for MockCreatures (not exported)
type _MockCreaturesRecorder struct {
	mock *MockCreatures
}

// NewMockCreatures function
func NewMockCreatures(ctrl *gomock.Controller) *MockCreatures {
	mock := &MockCreatures{ctrl: ctrl}
	mock.recorder = &_MockCreaturesRecorder{mock}
	return mock
}

// EXPECT function
func (_m *MockCreatures) EXPECT() *_MockCreaturesRecorder {
	return _m.recorder
}

// Kind function
func (_m *MockCreatures) Kind() string {
	ret := _m.ctrl.Call(_m, "Kind")
	ret0, _ := ret[0].(string)
	return ret0
}

// MockCreaturesRecorder Kind
func (_mr *_MockCreaturesRecorder) Kind() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Kind")
}

// Fly function
func (_m *MockCreatures) Fly() bool {
	ret := _m.ctrl.Call(_m, "Fly")
	ret0, _ := ret[0].(bool)
	return ret0
}

// MockCreaturesRecorder Fly
func (_mr *_MockCreaturesRecorder) Fly() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fly")
}

// Sound function
func (_m *MockCreatures) Sound() string {
	ret := _m.ctrl.Call(_m, "Sound")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockCreaturesRecorder) Sound() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Sound")
}
