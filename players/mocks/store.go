// Code generated by MockGen. DO NOT EDIT.
// Source: ./spec.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	model "surasithaof/gin-graphql-server/players/model"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStore) Create(ctx context.Context, team *model.Player) (*model.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, team)
	ret0, _ := ret[0].(*model.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStoreMockRecorder) Create(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), ctx, team)
}

// Delete mocks base method.
func (m *MockStore) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStoreMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), ctx, id)
}

// FindAll mocks base method.
func (m *MockStore) FindAll(ctx context.Context) ([]*model.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]*model.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockStoreMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockStore)(nil).FindAll), ctx)
}

// FindByIDs mocks base method.
func (m *MockStore) FindByIDs(ctx context.Context, IDs []string) (map[string]*model.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDs", ctx, IDs)
	ret0, _ := ret[0].(map[string]*model.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDs indicates an expected call of FindByIDs.
func (mr *MockStoreMockRecorder) FindByIDs(ctx, IDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDs", reflect.TypeOf((*MockStore)(nil).FindByIDs), ctx, IDs)
}

// FindByTeamId mocks base method.
func (m *MockStore) FindByTeamId(ctx context.Context, teamId int) ([]*model.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTeamId", ctx, teamId)
	ret0, _ := ret[0].([]*model.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTeamId indicates an expected call of FindByTeamId.
func (mr *MockStoreMockRecorder) FindByTeamId(ctx, teamId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTeamId", reflect.TypeOf((*MockStore)(nil).FindByTeamId), ctx, teamId)
}

// FindByTeamIds mocks base method.
func (m *MockStore) FindByTeamIds(ctx context.Context, teamIds []int) ([]*model.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTeamIds", ctx, teamIds)
	ret0, _ := ret[0].([]*model.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTeamIds indicates an expected call of FindByTeamIds.
func (mr *MockStoreMockRecorder) FindByTeamIds(ctx, teamIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTeamIds", reflect.TypeOf((*MockStore)(nil).FindByTeamIds), ctx, teamIds)
}

// FindOne mocks base method.
func (m *MockStore) FindOne(ctx context.Context, id int) (*model.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx, id)
	ret0, _ := ret[0].(*model.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockStoreMockRecorder) FindOne(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockStore)(nil).FindOne), ctx, id)
}

// Update mocks base method.
func (m *MockStore) Update(ctx context.Context, id int, team *model.Player) (*model.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, team)
	ret0, _ := ret[0].(*model.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockStoreMockRecorder) Update(ctx, id, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), ctx, id, team)
}
