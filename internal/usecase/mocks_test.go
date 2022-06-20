// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	entity "devops-tpl/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDevOps is a mock of DevOps interface.
type MockDevOps struct {
	ctrl     *gomock.Controller
	recorder *MockDevOpsMockRecorder
}

// MockDevOpsMockRecorder is the mock recorder for MockDevOps.
type MockDevOpsMockRecorder struct {
	mock *MockDevOps
}

// NewMockDevOps creates a new mock instance.
func NewMockDevOps(ctrl *gomock.Controller) *MockDevOps {
	mock := &MockDevOps{ctrl: ctrl}
	mock.recorder = &MockDevOpsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDevOps) EXPECT() *MockDevOpsMockRecorder {
	return m.recorder
}

// Metric mocks base method.
func (m *MockDevOps) Metric(arg0 context.Context, arg1 entity.Metric) (entity.Metric, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metric", arg0, arg1)
	ret0, _ := ret[0].(entity.Metric)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metric indicates an expected call of Metric.
func (mr *MockDevOpsMockRecorder) Metric(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metric", reflect.TypeOf((*MockDevOps)(nil).Metric), arg0, arg1)
}

// MetricNames mocks base method.
func (m *MockDevOps) MetricNames(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MetricNames", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MetricNames indicates an expected call of MetricNames.
func (mr *MockDevOpsMockRecorder) MetricNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MetricNames", reflect.TypeOf((*MockDevOps)(nil).MetricNames), arg0)
}

// StoreMetric mocks base method.
func (m *MockDevOps) StoreMetric(arg0 context.Context, arg1 entity.Metric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreMetric", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreMetric indicates an expected call of StoreMetric.
func (mr *MockDevOpsMockRecorder) StoreMetric(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreMetric", reflect.TypeOf((*MockDevOps)(nil).StoreMetric), arg0, arg1)
}

// saveStorage mocks base method.
func (m *MockDevOps) saveStorage() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "saveStorage")
}

// saveStorage indicates an expected call of saveStorage.
func (mr *MockDevOpsMockRecorder) saveStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "saveStorage", reflect.TypeOf((*MockDevOps)(nil).saveStorage))
}

// MockMetricRepo is a mock of MetricRepo interface.
type MockMetricRepo struct {
	ctrl     *gomock.Controller
	recorder *MockMetricRepoMockRecorder
}

// MockMetricRepoMockRecorder is the mock recorder for MockMetricRepo.
type MockMetricRepoMockRecorder struct {
	mock *MockMetricRepo
}

// NewMockMetricRepo creates a new mock instance.
func NewMockMetricRepo(ctrl *gomock.Controller) *MockMetricRepo {
	mock := &MockMetricRepo{ctrl: ctrl}
	mock.recorder = &MockMetricRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricRepo) EXPECT() *MockMetricRepoMockRecorder {
	return m.recorder
}

// GetMetric mocks base method.
func (m *MockMetricRepo) GetMetric(arg0 context.Context, arg1 string) (entity.Metric, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetric", arg0, arg1)
	ret0, _ := ret[0].(entity.Metric)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetric indicates an expected call of GetMetric.
func (mr *MockMetricRepoMockRecorder) GetMetric(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetric", reflect.TypeOf((*MockMetricRepo)(nil).GetMetric), arg0, arg1)
}

// GetMetricNames mocks base method.
func (m *MockMetricRepo) GetMetricNames(arg0 context.Context) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetricNames", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetMetricNames indicates an expected call of GetMetricNames.
func (mr *MockMetricRepoMockRecorder) GetMetricNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricNames", reflect.TypeOf((*MockMetricRepo)(nil).GetMetricNames), arg0)
}

// StoreMetric mocks base method.
func (m *MockMetricRepo) StoreMetric(arg0 context.Context, arg1 entity.Metric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreMetric", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreMetric indicates an expected call of StoreMetric.
func (mr *MockMetricRepoMockRecorder) StoreMetric(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreMetric", reflect.TypeOf((*MockMetricRepo)(nil).StoreMetric), arg0, arg1)
}

// StoreToFile mocks base method.
func (m *MockMetricRepo) StoreToFile() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreToFile")
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreToFile indicates an expected call of StoreToFile.
func (mr *MockMetricRepoMockRecorder) StoreToFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreToFile", reflect.TypeOf((*MockMetricRepo)(nil).StoreToFile))
}

// UploadFromFile mocks base method.
func (m *MockMetricRepo) UploadFromFile(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFromFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFromFile indicates an expected call of UploadFromFile.
func (mr *MockMetricRepoMockRecorder) UploadFromFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFromFile", reflect.TypeOf((*MockMetricRepo)(nil).UploadFromFile), arg0)
}