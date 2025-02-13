// Code generated by MockGen. DO NOT EDIT.
// Source: registry_client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// MockRegistryClient is a mock of RegistryClient interface.
type MockRegistryClient struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryClientMockRecorder
}

// MockRegistryClientMockRecorder is the mock recorder for MockRegistryClient.
type MockRegistryClientMockRecorder struct {
	mock *MockRegistryClient
}

// NewMockRegistryClient creates a new mock instance.
func NewMockRegistryClient(ctrl *gomock.Controller) *MockRegistryClient {
	mock := &MockRegistryClient{ctrl: ctrl}
	mock.recorder = &MockRegistryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistryClient) EXPECT() *MockRegistryClientMockRecorder {
	return m.recorder
}

// DownloadBundle mocks base method.
func (m *MockRegistryClient) DownloadBundle(ctx context.Context, ref string) (*v1alpha1.PackageBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadBundle", ctx, ref)
	ret0, _ := ret[0].(*v1alpha1.PackageBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadBundle indicates an expected call of DownloadBundle.
func (mr *MockRegistryClientMockRecorder) DownloadBundle(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadBundle", reflect.TypeOf((*MockRegistryClient)(nil).DownloadBundle), ctx, ref)
}

// LatestBundle mocks base method.
func (m *MockRegistryClient) LatestBundle(ctx context.Context, baseRef, kubeMajor, kubeMinor string) (*v1alpha1.PackageBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestBundle", ctx, baseRef, kubeMajor, kubeMinor)
	ret0, _ := ret[0].(*v1alpha1.PackageBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestBundle indicates an expected call of LatestBundle.
func (mr *MockRegistryClientMockRecorder) LatestBundle(ctx, baseRef, kubeMajor, kubeMinor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestBundle", reflect.TypeOf((*MockRegistryClient)(nil).LatestBundle), ctx, baseRef, kubeMajor, kubeMinor)
}
