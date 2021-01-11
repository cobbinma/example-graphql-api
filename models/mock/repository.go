// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	models "github.com/cobbinma/example-graphql-api/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// MenuItems mocks base method
func (m *MockRepository) MenuItems(ctx context.Context) ([]*models.MenuItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MenuItems", ctx)
	ret0, _ := ret[0].([]*models.MenuItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MenuItems indicates an expected call of MenuItems
func (mr *MockRepositoryMockRecorder) MenuItems(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MenuItems", reflect.TypeOf((*MockRepository)(nil).MenuItems), ctx)
}

// UpdateMenuItems mocks base method
func (m *MockRepository) UpdateMenuItems(ctx context.Context, items []*models.MenuItemInput) ([]*models.MenuItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMenuItems", ctx, items)
	ret0, _ := ret[0].([]*models.MenuItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMenuItems indicates an expected call of UpdateMenuItems
func (mr *MockRepositoryMockRecorder) UpdateMenuItems(ctx, items interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMenuItems", reflect.TypeOf((*MockRepository)(nil).UpdateMenuItems), ctx, items)
}
