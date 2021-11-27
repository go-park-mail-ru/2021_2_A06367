// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package actors is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockActorsUsecase is a mock of ActorsUsecase interface.
type MockActorsUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockActorsUsecaseMockRecorder
}

// MockActorsUsecaseMockRecorder is the mock recorder for MockActorsUsecase.
type MockActorsUsecaseMockRecorder struct {
	mock *MockActorsUsecase
}

// NewMockActorsUsecase creates a new mock instance.
func NewMockActorsUsecase(ctrl *gomock.Controller) *MockActorsUsecase {
	mock := &MockActorsUsecase{ctrl: ctrl}
	mock.recorder = &MockActorsUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActorsUsecase) EXPECT() *MockActorsUsecaseMockRecorder {
	return m.recorder
}

// GetByActors mocks base method.
func (m *MockActorsUsecase) GetByActors(actor []models.Actors) ([]models.Actors, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByActors", actor)
	ret0, _ := ret[0].([]models.Actors)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// GetByActors indicates an expected call of GetByActors.
func (mr *MockActorsUsecaseMockRecorder) GetByActors(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByActors", reflect.TypeOf((*MockActorsUsecase)(nil).GetByActors), actor)
}

// GetById mocks base method.
func (m *MockActorsUsecase) GetById(actor models.Actors) (models.Actors, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", actor)
	ret0, _ := ret[0].(models.Actors)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockActorsUsecaseMockRecorder) GetById(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockActorsUsecase)(nil).GetById), actor)
}

// GetByKeyword mocks base method.
func (m *MockActorsUsecase) GetByKeyword(keyword string) ([]models.Actors, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByKeyword", keyword)
	ret0, _ := ret[0].([]models.Actors)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// GetByKeyword indicates an expected call of GetByKeyword.
func (mr *MockActorsUsecaseMockRecorder) GetByKeyword(keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByKeyword", reflect.TypeOf((*MockActorsUsecase)(nil).GetByKeyword), keyword)
}

// MockActorsRepository is a mock of ActorsRepository interface.
type MockActorsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockActorsRepositoryMockRecorder
}

// MockActorsRepositoryMockRecorder is the mock recorder for MockActorsRepository.
type MockActorsRepositoryMockRecorder struct {
	mock *MockActorsRepository
}

// NewMockActorsRepository creates a new mock instance.
func NewMockActorsRepository(ctrl *gomock.Controller) *MockActorsRepository {
	mock := &MockActorsRepository{ctrl: ctrl}
	mock.recorder = &MockActorsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActorsRepository) EXPECT() *MockActorsRepositoryMockRecorder {
	return m.recorder
}

// GetActorById mocks base method.
func (m *MockActorsRepository) GetActorById(actor models.Actors) (models.Actors, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorById", actor)
	ret0, _ := ret[0].(models.Actors)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// GetActorById indicates an expected call of GetActorById.
func (mr *MockActorsRepositoryMockRecorder) GetActorById(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorById", reflect.TypeOf((*MockActorsRepository)(nil).GetActorById), actor)
}

// GetActors mocks base method.
func (m *MockActorsRepository) GetActors(actor []models.Actors) ([]models.Actors, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActors", actor)
	ret0, _ := ret[0].([]models.Actors)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// GetActors indicates an expected call of GetActors.
func (mr *MockActorsRepositoryMockRecorder) GetActors(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActors", reflect.TypeOf((*MockActorsRepository)(nil).GetActors), actor)
}

// GetActorsByKeyword mocks base method.
func (m *MockActorsRepository) GetActorsByKeyword(keyword string) ([]models.Actors, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorsByKeyword", keyword)
	ret0, _ := ret[0].([]models.Actors)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// GetActorsByKeyword indicates an expected call of GetActorsByKeyword.
func (mr *MockActorsRepositoryMockRecorder) GetActorsByKeyword(keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorsByKeyword", reflect.TypeOf((*MockActorsRepository)(nil).GetActorsByKeyword), keyword)
}