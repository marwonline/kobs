// Code generated by MockGen. DO NOT EDIT.
// Source: instance.go

// Package instance is a generated GoMock package.
package instance

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	model "github.com/prometheus/common/model"
)

// MockInstance is a mock of Instance interface.
type MockInstance struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceMockRecorder
}

// MockInstanceMockRecorder is the mock recorder for MockInstance.
type MockInstanceMockRecorder struct {
	mock *MockInstance
}

// NewMockInstance creates a new mock instance.
func NewMockInstance(ctrl *gomock.Controller) *MockInstance {
	mock := &MockInstance{ctrl: ctrl}
	mock.recorder = &MockInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstance) EXPECT() *MockInstanceMockRecorder {
	return m.recorder
}

// GetInstant mocks base method.
func (m *MockInstance) GetInstant(ctx context.Context, queries []Query, timeEnd int64) (map[string]map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstant", ctx, queries, timeEnd)
	ret0, _ := ret[0].(map[string]map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstant indicates an expected call of GetInstant.
func (mr *MockInstanceMockRecorder) GetInstant(ctx, queries, timeEnd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstant", reflect.TypeOf((*MockInstance)(nil).GetInstant), ctx, queries, timeEnd)
}

// GetLabels mocks base method.
func (m *MockInstance) GetLabels(ctx context.Context, matches []string, timeStart, timeEnd int64) ([]string, v1.Warnings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabels", ctx, matches, timeStart, timeEnd)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(v1.Warnings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLabels indicates an expected call of GetLabels.
func (mr *MockInstanceMockRecorder) GetLabels(ctx, matches, timeStart, timeEnd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockInstance)(nil).GetLabels), ctx, matches, timeStart, timeEnd)
}

// GetName mocks base method.
func (m *MockInstance) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockInstanceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockInstance)(nil).GetName))
}

// GetRange mocks base method.
func (m *MockInstance) GetRange(ctx context.Context, queries []Query, resolution string, timeStart, timeEnd int64) (*Metrics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRange", ctx, queries, resolution, timeStart, timeEnd)
	ret0, _ := ret[0].(*Metrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRange indicates an expected call of GetRange.
func (mr *MockInstanceMockRecorder) GetRange(ctx, queries, resolution, timeStart, timeEnd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRange", reflect.TypeOf((*MockInstance)(nil).GetRange), ctx, queries, resolution, timeStart, timeEnd)
}

// GetSeries mocks base method.
func (m *MockInstance) GetSeries(ctx context.Context, matches []string, timeStart, timeEnd int64) ([]model.LabelSet, v1.Warnings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeries", ctx, matches, timeStart, timeEnd)
	ret0, _ := ret[0].([]model.LabelSet)
	ret1, _ := ret[1].(v1.Warnings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSeries indicates an expected call of GetSeries.
func (mr *MockInstanceMockRecorder) GetSeries(ctx, matches, timeStart, timeEnd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeries", reflect.TypeOf((*MockInstance)(nil).GetSeries), ctx, matches, timeStart, timeEnd)
}

// GetVariable mocks base method.
func (m *MockInstance) GetVariable(ctx context.Context, label, query, queryType string, timeStart, timeEnd int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVariable", ctx, label, query, queryType, timeStart, timeEnd)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVariable indicates an expected call of GetVariable.
func (mr *MockInstanceMockRecorder) GetVariable(ctx, label, query, queryType, timeStart, timeEnd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVariable", reflect.TypeOf((*MockInstance)(nil).GetVariable), ctx, label, query, queryType, timeStart, timeEnd)
}
