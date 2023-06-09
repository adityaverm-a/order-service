// Code generated by MockGen. DO NOT EDIT.
// Source: demo/oms/order/domain/repositories (interfaces: OrderRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "demo/oms/order/data/models"
	entities "demo/oms/order/domain/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderRepository) Create(arg0 entities.CreateOrderInput) (*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderRepository)(nil).Create), arg0)
}

// GetByFilters mocks base method.
func (m *MockOrderRepository) GetByFilters(arg0 entities.OrderFiltersInput) (*[]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByFilters", arg0)
	ret0, _ := ret[0].(*[]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByFilters indicates an expected call of GetByFilters.
func (mr *MockOrderRepositoryMockRecorder) GetByFilters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByFilters", reflect.TypeOf((*MockOrderRepository)(nil).GetByFilters), arg0)
}

// GetByID mocks base method.
func (m *MockOrderRepository) GetByID(arg0 int64) (*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockOrderRepositoryMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockOrderRepository)(nil).GetByID), arg0)
}

// Update mocks base method.
func (m *MockOrderRepository) Update(arg0 entities.UpdateOrderStatusInput) (*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrderRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderRepository)(nil).Update), arg0)
}
