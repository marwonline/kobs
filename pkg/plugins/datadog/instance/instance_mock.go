// Code generated by MockGen. DO NOT EDIT.
// Source: instance.go

// Package instance is a generated GoMock package.
package instance

import (
	context "context"
	reflect "reflect"

	datadogV1 "github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	datadogV2 "github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	gomock "github.com/golang/mock/gomock"
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

// GetLogs mocks base method.
func (m *MockInstance) GetLogs(ctx context.Context, query string, startTime, endTime int64) ([]datadogV2.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", ctx, query, startTime, endTime)
	ret0, _ := ret[0].([]datadogV2.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockInstanceMockRecorder) GetLogs(ctx, query, startTime, endTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockInstance)(nil).GetLogs), ctx, query, startTime, endTime)
}

// GetLogsAggregation mocks base method.
func (m *MockInstance) GetLogsAggregation(ctx context.Context, query string, startTime, endTime int64) ([]datadogV2.LogsAggregateBucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogsAggregation", ctx, query, startTime, endTime)
	ret0, _ := ret[0].([]datadogV2.LogsAggregateBucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogsAggregation indicates an expected call of GetLogsAggregation.
func (mr *MockInstanceMockRecorder) GetLogsAggregation(ctx, query, startTime, endTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogsAggregation", reflect.TypeOf((*MockInstance)(nil).GetLogsAggregation), ctx, query, startTime, endTime)
}

// GetMetrics mocks base method.
func (m *MockInstance) GetMetrics(ctx context.Context, query string, startTime, endTime int64) (*datadogV1.MetricsQueryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetrics", ctx, query, startTime, endTime)
	ret0, _ := ret[0].(*datadogV1.MetricsQueryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetrics indicates an expected call of GetMetrics.
func (mr *MockInstanceMockRecorder) GetMetrics(ctx, query, startTime, endTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetrics", reflect.TypeOf((*MockInstance)(nil).GetMetrics), ctx, query, startTime, endTime)
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