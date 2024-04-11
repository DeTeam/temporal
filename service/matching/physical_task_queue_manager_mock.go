// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: service/matching/physical_task_queue_manager.go

// Package matching is a generated GoMock package.
package matching

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1 "go.temporal.io/api/taskqueue/v1"
	v10 "go.temporal.io/server/api/matchingservice/v1"
	v11 "go.temporal.io/server/api/taskqueue/v1"
)

// MockphysicalTaskQueueManager is a mock of physicalTaskQueueManager interface.
type MockphysicalTaskQueueManager struct {
	ctrl     *gomock.Controller
	recorder *MockphysicalTaskQueueManagerMockRecorder
}

// MockphysicalTaskQueueManagerMockRecorder is the mock recorder for MockphysicalTaskQueueManager.
type MockphysicalTaskQueueManagerMockRecorder struct {
	mock *MockphysicalTaskQueueManager
}

// NewMockphysicalTaskQueueManager creates a new mock instance.
func NewMockphysicalTaskQueueManager(ctrl *gomock.Controller) *MockphysicalTaskQueueManager {
	mock := &MockphysicalTaskQueueManager{ctrl: ctrl}
	mock.recorder = &MockphysicalTaskQueueManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockphysicalTaskQueueManager) EXPECT() *MockphysicalTaskQueueManagerMockRecorder {
	return m.recorder
}

// AddTask mocks base method.
func (m *MockphysicalTaskQueueManager) AddTask(ctx context.Context, params addTaskParams) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", ctx, params)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTask indicates an expected call of AddTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) AddTask(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).AddTask), ctx, params)
}

// Describe mocks base method.
func (m *MockphysicalTaskQueueManager) Describe() *v11.PhysicalTaskQueueInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Describe")
	ret0, _ := ret[0].(*v11.PhysicalTaskQueueInfo)
	return ret0
}

// Describe indicates an expected call of Describe.
func (mr *MockphysicalTaskQueueManagerMockRecorder) Describe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).Describe))
}

// DispatchQueryTask mocks base method.
func (m *MockphysicalTaskQueueManager) DispatchQueryTask(ctx context.Context, taskID string, request *v10.QueryWorkflowRequest) (*v10.QueryWorkflowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchQueryTask", ctx, taskID, request)
	ret0, _ := ret[0].(*v10.QueryWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DispatchQueryTask indicates an expected call of DispatchQueryTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) DispatchQueryTask(ctx, taskID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchQueryTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).DispatchQueryTask), ctx, taskID, request)
}

// DispatchSpooledTask mocks base method.
func (m *MockphysicalTaskQueueManager) DispatchSpooledTask(ctx context.Context, task *internalTask, userDataChanged <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchSpooledTask", ctx, task, userDataChanged)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchSpooledTask indicates an expected call of DispatchSpooledTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) DispatchSpooledTask(ctx, task, userDataChanged interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchSpooledTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).DispatchSpooledTask), ctx, task, userDataChanged)
}

// GetAllPollerInfo mocks base method.
func (m *MockphysicalTaskQueueManager) GetAllPollerInfo() []*v1.PollerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPollerInfo")
	ret0, _ := ret[0].([]*v1.PollerInfo)
	return ret0
}

// GetAllPollerInfo indicates an expected call of GetAllPollerInfo.
func (mr *MockphysicalTaskQueueManagerMockRecorder) GetAllPollerInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPollerInfo", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).GetAllPollerInfo))
}

// GetBacklogInfo mocks base method.
func (m *MockphysicalTaskQueueManager) GetBacklogInfo() *v1.BacklogInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBacklogInfo")
	ret0, _ := ret[0].(*v1.BacklogInfo)
	return ret0
}

// GetBacklogInfo indicates an expected call of GetBacklogInfo.
func (mr *MockphysicalTaskQueueManagerMockRecorder) GetBacklogInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBacklogInfo", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).GetBacklogInfo))
}

// HasPollerAfter mocks base method.
func (m *MockphysicalTaskQueueManager) HasPollerAfter(accessTime time.Time) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPollerAfter", accessTime)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPollerAfter indicates an expected call of HasPollerAfter.
func (mr *MockphysicalTaskQueueManagerMockRecorder) HasPollerAfter(accessTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPollerAfter", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).HasPollerAfter), accessTime)
}

// LegacyDescribeTaskQueue mocks base method.
func (m *MockphysicalTaskQueueManager) LegacyDescribeTaskQueue(includeTaskQueueStatus bool) *v10.DescribeTaskQueueResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LegacyDescribeTaskQueue", includeTaskQueueStatus)
	ret0, _ := ret[0].(*v10.DescribeTaskQueueResponse)
	return ret0
}

// LegacyDescribeTaskQueue indicates an expected call of LegacyDescribeTaskQueue.
func (mr *MockphysicalTaskQueueManagerMockRecorder) LegacyDescribeTaskQueue(includeTaskQueueStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LegacyDescribeTaskQueue", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).LegacyDescribeTaskQueue), includeTaskQueueStatus)
}

// MarkAlive mocks base method.
func (m *MockphysicalTaskQueueManager) MarkAlive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkAlive")
}

// MarkAlive indicates an expected call of MarkAlive.
func (mr *MockphysicalTaskQueueManagerMockRecorder) MarkAlive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAlive", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).MarkAlive))
}

// PollTask mocks base method.
func (m *MockphysicalTaskQueueManager) PollTask(ctx context.Context, pollMetadata *pollMetadata) (*internalTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollTask", ctx, pollMetadata)
	ret0, _ := ret[0].(*internalTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollTask indicates an expected call of PollTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) PollTask(ctx, pollMetadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).PollTask), ctx, pollMetadata)
}

// ProcessSpooledTask mocks base method.
func (m *MockphysicalTaskQueueManager) ProcessSpooledTask(ctx context.Context, task *internalTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessSpooledTask", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessSpooledTask indicates an expected call of ProcessSpooledTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) ProcessSpooledTask(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSpooledTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).ProcessSpooledTask), ctx, task)
}

// QueueKey mocks base method.
func (m *MockphysicalTaskQueueManager) QueueKey() *PhysicalTaskQueueKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueKey")
	ret0, _ := ret[0].(*PhysicalTaskQueueKey)
	return ret0
}

// QueueKey indicates an expected call of QueueKey.
func (mr *MockphysicalTaskQueueManagerMockRecorder) QueueKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueKey", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).QueueKey))
}

// SpoolTask mocks base method.
func (m *MockphysicalTaskQueueManager) SpoolTask(params addTaskParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpoolTask", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// SpoolTask indicates an expected call of SpoolTask.
func (mr *MockphysicalTaskQueueManagerMockRecorder) SpoolTask(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpoolTask", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).SpoolTask), params)
}

// Start mocks base method.
func (m *MockphysicalTaskQueueManager) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockphysicalTaskQueueManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).Start))
}

// Stop mocks base method.
func (m *MockphysicalTaskQueueManager) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockphysicalTaskQueueManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).Stop))
}

// String mocks base method.
func (m *MockphysicalTaskQueueManager) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockphysicalTaskQueueManagerMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).String))
}

// UnloadFromPartitionManager mocks base method.
func (m *MockphysicalTaskQueueManager) UnloadFromPartitionManager() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnloadFromPartitionManager")
}

// UnloadFromPartitionManager indicates an expected call of UnloadFromPartitionManager.
func (mr *MockphysicalTaskQueueManagerMockRecorder) UnloadFromPartitionManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnloadFromPartitionManager", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).UnloadFromPartitionManager))
}

// UpdatePollerInfo mocks base method.
func (m *MockphysicalTaskQueueManager) UpdatePollerInfo(arg0 pollerIdentity, arg1 *pollMetadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePollerInfo", arg0, arg1)
}

// UpdatePollerInfo indicates an expected call of UpdatePollerInfo.
func (mr *MockphysicalTaskQueueManagerMockRecorder) UpdatePollerInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePollerInfo", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).UpdatePollerInfo), arg0, arg1)
}

// WaitUntilInitialized mocks base method.
func (m *MockphysicalTaskQueueManager) WaitUntilInitialized(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilInitialized", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilInitialized indicates an expected call of WaitUntilInitialized.
func (mr *MockphysicalTaskQueueManagerMockRecorder) WaitUntilInitialized(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilInitialized", reflect.TypeOf((*MockphysicalTaskQueueManager)(nil).WaitUntilInitialized), arg0)
}
