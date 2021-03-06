// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package auth is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockAuthUsecase is a mock of AuthUsecase interface.
type MockAuthUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUsecaseMockRecorder
}

// MockAuthUsecaseMockRecorder is the mock recorder for MockAuthUsecase.
type MockAuthUsecaseMockRecorder struct {
	mock *MockAuthUsecase
}

// NewMockAuthUsecase creates a new mock instance.
func NewMockAuthUsecase(ctrl *gomock.Controller) *MockAuthUsecase {
	mock := &MockAuthUsecase{ctrl: ctrl}
	mock.recorder = &MockAuthUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUsecase) EXPECT() *MockAuthUsecaseMockRecorder {
	return m.recorder
}

// SignIn mocks base method.
func (m *MockAuthUsecase) SignIn(user models.LoginUser) (string, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockAuthUsecaseMockRecorder) SignIn(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAuthUsecase)(nil).SignIn), user)
}

// SignUp mocks base method.
func (m *MockAuthUsecase) SignUp(user models.User) (string, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockAuthUsecaseMockRecorder) SignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthUsecase)(nil).SignUp), user)
}

// GetProfile mocks base method.
func (m *MockAuthUsecase) GetProfile(user models.Profile) (models.Profile, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", user)
	ret0, _ := ret[0].(models.Profile)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockAuthUsecaseMockRecorder) GetProfile(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockAuthUsecase)(nil).GetProfile), user)
}

// Follow mocks base method.
func (m *MockAuthUsecase) Follow(who, whom uuid.UUID) models.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", who, whom)
	ret0, _ := ret[0].(models.StatusCode)
	return ret0
}

// Follow indicates an expected call of Follow.
func (mr *MockAuthUsecaseMockRecorder) Follow(who, whom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockAuthUsecase)(nil).Follow), who, whom)
}

// GetByKeyword mocks base method.
func (m *MockAuthUsecase) GetByKeyword(keyword string) ([]models.Profile, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByKeyword", keyword)
	ret0, _ := ret[0].([]models.Profile)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// GetByKeyword indicates an expected call of GetByKeyword.
func (mr *MockAuthUsecaseMockRecorder) GetByKeyword(keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByKeyword", reflect.TypeOf((*MockAuthUsecase)(nil).GetByKeyword), keyword)
}

// SetBio mocks base method.
func (m *MockAuthUsecase) SetBio(profile models.Profile) models.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBio", profile)
	ret0, _ := ret[0].(models.StatusCode)
	return ret0
}

// SetBio indicates an expected call of SetBio.
func (mr *MockAuthUsecaseMockRecorder) SetBio(profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBio", reflect.TypeOf((*MockAuthUsecase)(nil).SetBio), profile)
}

// SetPass mocks base method.
func (m *MockAuthUsecase) SetPass(profile models.User) models.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPass", profile)
	ret0, _ := ret[0].(models.StatusCode)
	return ret0
}

// SetPass indicates an expected call of SetPass.
func (mr *MockAuthUsecaseMockRecorder) SetPass(profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPass", reflect.TypeOf((*MockAuthUsecase)(nil).SetPass), profile)
}

// SetAvatar mocks base method.
func (m *MockAuthUsecase) SetAvatar(profile models.Profile) models.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAvatar", profile)
	ret0, _ := ret[0].(models.StatusCode)
	return ret0
}

// SetAvatar indicates an expected call of SetAvatar.
func (mr *MockAuthUsecaseMockRecorder) SetAvatar(profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAvatar", reflect.TypeOf((*MockAuthUsecase)(nil).SetAvatar), profile)
}

// CheckUser mocks base method.
func (m *MockAuthUsecase) CheckUser(user models.User) (models.User, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUser", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// CheckUser indicates an expected call of CheckUser.
func (mr *MockAuthUsecaseMockRecorder) CheckUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUser", reflect.TypeOf((*MockAuthUsecase)(nil).CheckUser), user)
}

// CheckUserLogin mocks base method.
func (m *MockAuthUsecase) CheckUserLogin(user models.User) (models.User, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserLogin", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// CheckUserLogin indicates an expected call of CheckUserLogin.
func (mr *MockAuthUsecaseMockRecorder) CheckUserLogin(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserLogin", reflect.TypeOf((*MockAuthUsecase)(nil).CheckUserLogin), user)
}

// MockAuthRepo is a mock of AuthRepo interface.
type MockAuthRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepoMockRecorder
}

// MockAuthRepoMockRecorder is the mock recorder for MockAuthRepo.
type MockAuthRepoMockRecorder struct {
	mock *MockAuthRepo
}

// NewMockAuthRepo creates a new mock instance.
func NewMockAuthRepo(ctrl *gomock.Controller) *MockAuthRepo {
	mock := &MockAuthRepo{ctrl: ctrl}
	mock.recorder = &MockAuthRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepo) EXPECT() *MockAuthRepoMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthRepo) CreateUser(user models.User) (models.User, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthRepoMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthRepo)(nil).CreateUser), user)
}

// CheckUser mocks base method.
func (m *MockAuthRepo) CheckUser(user models.User) (models.User, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUser", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// CheckUser indicates an expected call of CheckUser.
func (mr *MockAuthRepoMockRecorder) CheckUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUser", reflect.TypeOf((*MockAuthRepo)(nil).CheckUser), user)
}

// CheckUserLogin mocks base method.
func (m *MockAuthRepo) CheckUserLogin(user models.User) (models.User, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserLogin", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// CheckUserLogin indicates an expected call of CheckUserLogin.
func (mr *MockAuthRepoMockRecorder) CheckUserLogin(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserLogin", reflect.TypeOf((*MockAuthRepo)(nil).CheckUserLogin), user)
}

// GetProfile mocks base method.
func (m *MockAuthRepo) GetProfile(user models.Profile) (models.Profile, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", user)
	ret0, _ := ret[0].(models.Profile)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockAuthRepoMockRecorder) GetProfile(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockAuthRepo)(nil).GetProfile), user)
}

// AddFollowing mocks base method.
func (m *MockAuthRepo) AddFollowing(who, whom uuid.UUID) models.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFollowing", who, whom)
	ret0, _ := ret[0].(models.StatusCode)
	return ret0
}

// AddFollowing indicates an expected call of AddFollowing.
func (mr *MockAuthRepoMockRecorder) AddFollowing(who, whom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFollowing", reflect.TypeOf((*MockAuthRepo)(nil).AddFollowing), who, whom)
}

// RemoveFollowing mocks base method.
func (m *MockAuthRepo) RemoveFollowing(who, whom uuid.UUID) models.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFollowing", who, whom)
	ret0, _ := ret[0].(models.StatusCode)
	return ret0
}

// RemoveFollowing indicates an expected call of RemoveFollowing.
func (mr *MockAuthRepoMockRecorder) RemoveFollowing(who, whom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFollowing", reflect.TypeOf((*MockAuthRepo)(nil).RemoveFollowing), who, whom)
}

// GetProfileByKeyword mocks base method.
func (m *MockAuthRepo) GetProfileByKeyword(keyword string) ([]models.Profile, models.StatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileByKeyword", keyword)
	ret0, _ := ret[0].([]models.Profile)
	ret1, _ := ret[1].(models.StatusCode)
	return ret0, ret1
}

// GetProfileByKeyword indicates an expected call of GetProfileByKeyword.
func (mr *MockAuthRepoMockRecorder) GetProfileByKeyword(keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileByKeyword", reflect.TypeOf((*MockAuthRepo)(nil).GetProfileByKeyword), keyword)
}

// UpdateBio mocks base method.
func (m *MockAuthRepo) UpdateBio(profile models.Profile) models.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBio", profile)
	ret0, _ := ret[0].(models.StatusCode)
	return ret0
}

// UpdateBio indicates an expected call of UpdateBio.
func (mr *MockAuthRepoMockRecorder) UpdateBio(profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBio", reflect.TypeOf((*MockAuthRepo)(nil).UpdateBio), profile)
}

// UpdatePass mocks base method.
func (m *MockAuthRepo) UpdatePass(profile models.User) models.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePass", profile)
	ret0, _ := ret[0].(models.StatusCode)
	return ret0
}

// UpdatePass indicates an expected call of UpdatePass.
func (mr *MockAuthRepoMockRecorder) UpdatePass(profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePass", reflect.TypeOf((*MockAuthRepo)(nil).UpdatePass), profile)
}

// UpdateAvatar mocks base method.
func (m *MockAuthRepo) UpdateAvatar(profile models.Profile) models.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", profile)
	ret0, _ := ret[0].(models.StatusCode)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockAuthRepoMockRecorder) UpdateAvatar(profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockAuthRepo)(nil).UpdateAvatar), profile)
}

// MockTokenGenerator is a mock of TokenGenerator interface.
type MockTokenGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenGeneratorMockRecorder
}

// MockTokenGeneratorMockRecorder is the mock recorder for MockTokenGenerator.
type MockTokenGeneratorMockRecorder struct {
	mock *MockTokenGenerator
}

// NewMockTokenGenerator creates a new mock instance.
func NewMockTokenGenerator(ctrl *gomock.Controller) *MockTokenGenerator {
	mock := &MockTokenGenerator{ctrl: ctrl}
	mock.recorder = &MockTokenGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenGenerator) EXPECT() *MockTokenGeneratorMockRecorder {
	return m.recorder
}

// GetToken mocks base method.
func (m *MockTokenGenerator) GetToken(user models.User) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", user)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetToken indicates an expected call of GetToken.
func (mr *MockTokenGeneratorMockRecorder) GetToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockTokenGenerator)(nil).GetToken), user)
}

// MockEncrypter is a mock of Encrypter interface.
type MockEncrypter struct {
	ctrl     *gomock.Controller
	recorder *MockEncrypterMockRecorder
}

// MockEncrypterMockRecorder is the mock recorder for MockEncrypter.
type MockEncrypterMockRecorder struct {
	mock *MockEncrypter
}

// NewMockEncrypter creates a new mock instance.
func NewMockEncrypter(ctrl *gomock.Controller) *MockEncrypter {
	mock := &MockEncrypter{ctrl: ctrl}
	mock.recorder = &MockEncrypterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncrypter) EXPECT() *MockEncrypterMockRecorder {
	return m.recorder
}

// EncryptPswd mocks base method.
func (m *MockEncrypter) EncryptPswd(pswd string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptPswd", pswd)
	ret0, _ := ret[0].(string)
	return ret0
}

// EncryptPswd indicates an expected call of EncryptPswd.
func (mr *MockEncrypterMockRecorder) EncryptPswd(pswd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptPswd", reflect.TypeOf((*MockEncrypter)(nil).EncryptPswd), pswd)
}
