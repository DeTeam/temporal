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
// Source: cache.go

// Package workflow is a generated GoMock package.
package workflow

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "go.temporal.io/api/common/v1"
	namespace "go.temporal.io/server/common/namespace"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// GetOrCreateWorkflowExecution mocks base method.
func (m *MockCache) GetOrCreateWorkflowExecution(ctx context.Context, namespaceID namespace.ID, execution v1.WorkflowExecution, caller CallerType) (Context, ReleaseCacheFunc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateWorkflowExecution", ctx, namespaceID, execution, caller)
	ret0, _ := ret[0].(Context)
	ret1, _ := ret[1].(ReleaseCacheFunc)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrCreateWorkflowExecution indicates an expected call of GetOrCreateWorkflowExecution.
func (mr *MockCacheMockRecorder) GetOrCreateWorkflowExecution(ctx, namespaceID, execution, caller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateWorkflowExecution", reflect.TypeOf((*MockCache)(nil).GetOrCreateWorkflowExecution), ctx, namespaceID, execution, caller)
}