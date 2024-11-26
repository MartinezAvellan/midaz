// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LerianStudio/midaz/components/ledger/internal/adapters/postgres/organization (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen --destination=organization.mock.go --package=organization . Repository
//

// Package organization is a generated GoMock package.
package organization

import (
	context "context"
	reflect "reflect"

	mmodel "github.com/LerianStudio/midaz/common/mmodel"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 context.Context, arg1 *mmodel.Organization) (*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockRepository) Delete(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockRepository) Find(arg0 context.Context, arg1 uuid.UUID) (*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepositoryMockRecorder) Find(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepository)(nil).Find), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockRepository) FindAll(arg0 context.Context, arg1, arg2 int) ([]*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll), arg0, arg1, arg2)
}

// ListByIDs mocks base method.
func (m *MockRepository) ListByIDs(arg0 context.Context, arg1 []uuid.UUID) ([]*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByIDs", arg0, arg1)
	ret0, _ := ret[0].([]*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByIDs indicates an expected call of ListByIDs.
func (mr *MockRepositoryMockRecorder) ListByIDs(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByIDs", reflect.TypeOf((*MockRepository)(nil).ListByIDs), arg0, arg1)
}

// Update mocks base method.
func (m *MockRepository) Update(arg0 context.Context, arg1 uuid.UUID, arg2 *mmodel.Organization) (*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), arg0, arg1, arg2)
}
