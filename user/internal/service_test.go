package user_test

import (
	"errors"

	"github.com/mkrashad/go-todo/user/internal"
	"github.com/mkrashad/go-todo/user/internal/mocks"

	// "github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"testing"
	// "gorm.io/gorm"
)

type UserServiceUnitTestSuite struct {
	suite.Suite
	underTest          user.Service
	mockUserRepository *mocks.Repository
}

var users = []user.User{
	{

		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@mail.com",
	},
	{
		FirstName: "Adam",
		LastName:  "Smith",
		Email:     "adam.smith@mail.com",
	},
}
var id uint64 = 1

func TestUserServiceUnitTestSuite(t *testing.T) {
	suite.Run(t, &UserServiceUnitTestSuite{})
}

func (ts *UserServiceUnitTestSuite) SetupSuite() {
	ts.mockUserRepository = new(mocks.Repository)
	ts.underTest = user.NewUserService(ts.mockUserRepository)
}

func (ts *UserServiceUnitTestSuite) TestGetAllUsers() {
	// given
	ts.mockUserRepository.On("GetAllUsers").Once().Return(users)
	// when
	result := ts.underTest.GetAllUsers()
	// then
	ts.Equal(users, result)
	ts.mockUserRepository.AssertExpectations(ts.T())
}

func (ts *UserServiceUnitTestSuite) TestGetUserById_ValidId() {
	// given

	ts.mockUserRepository.On("GetUserById", id).Once().Return(users[0], nil)
	// when
	result, err := ts.underTest.GetUserById(id)
	// then
	ts.Equal(users[0], result)
	ts.NoError(err)
	ts.mockUserRepository.AssertExpectations(ts.T())
}

func (ts *UserServiceUnitTestSuite) TestGetUserById_InvalidId() {
	// given
	var id uint64 = 9999
	ts.mockUserRepository.On("GetUserById", id).Once().Return(user.User{}, errors.New("not found"))
	// when
	result, err := ts.underTest.GetUserById(id)
	// then
	ts.Zero(result)
	ts.Error(err)
	ts.mockUserRepository.AssertExpectations(ts.T())
}

func (ts *UserServiceUnitTestSuite) TestCreateUserValid() {
	// given
	newUser := user.User{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane.doe@mail.com",
	}
	ts.mockUserRepository.On("CreateUser", newUser).Once().Return(newUser, nil)
	// when
	result, err := ts.underTest.CreateUser(newUser)
	// then
	ts.Equal(newUser, result)
	ts.NoError(err)
	ts.mockUserRepository.AssertExpectations(ts.T())
}

func (ts *UserServiceUnitTestSuite) TestCreateUserInValid() {
	// given
	newUser := user.User{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane.doe@mail.com",
	}
	ts.mockUserRepository.On("CreateUser", newUser).Once().Return(user.User{}, errors.New("error"))
	// when
	_, err := ts.underTest.CreateUser(newUser)

	// then
	ts.Error(err)
	ts.mockUserRepository.AssertExpectations(ts.T())
}

func (ts *UserServiceUnitTestSuite) TestUpdateUserById_ValidUpdate() {
	// given
	updatedUser := users[0]
	updatedUser.FirstName = "Edward"
	ts.mockUserRepository.On("UpdateUserById", id, updatedUser).Once().Return(updatedUser, nil)
	// when
	result, err := ts.underTest.UpdateUserById(id, updatedUser)

	// then
	ts.NoError(err)
	ts.Equal(updatedUser, result)
	ts.mockUserRepository.AssertExpectations(ts.T())
}

func (ts *UserServiceUnitTestSuite) TestUpdateUserById_InvalidUpdate() {
	// given
	updatedUser := users[0]
	ts.mockUserRepository.On("UpdateUserById", id, updatedUser).Once().Return(user.User{}, errors.New("error"))
	// when
	result, err := ts.underTest.UpdateUserById(id, updatedUser)
	// then
	ts.Error(err)
	ts.Zero(result)
	ts.mockUserRepository.AssertExpectations(ts.T())
}

func (ts *UserServiceUnitTestSuite) TestDeleteUserById_Valid() {
	// given
	ts.mockUserRepository.On("DeleteUserById", id).Once().Return(nil)
	// when
	ts.underTest.DeleteUserById(id)
	// then
	ts.mockUserRepository.AssertExpectations(ts.T())
}

func (ts *UserServiceUnitTestSuite) TestDeleteUserById_Invalid() {
	// given
	ts.mockUserRepository.On("DeleteUserById", id).Once().Return(errors.New("error"))

	// when
	err := ts.underTest.DeleteUserById(id)

	// then
	ts.Error(err)
	ts.mockUserRepository.AssertExpectations(ts.T())

}

func BenchmarkTaskService_GetAllUsers(b *testing.B) {
	mockUserRepository := new(mocks.Repository)
	underTest := user.NewUserService(mockUserRepository)

	mockUserRepository.On("GetAllUsers").Return(users)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		underTest.GetAllUsers()
	}
}

func BenchmarkTaskService_GetUserById(b *testing.B) {
	mockUserRepository := new(mocks.Repository)
	underTest := user.NewUserService(mockUserRepository)

	mockUserRepository.On("GetUserById", id).Return(users[0], nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = underTest.GetUserById(id)
	}
}

func BenchmarkTaskService_CreateUser(b *testing.B) {
	mockUserRepository := new(mocks.Repository)
	underTest := user.NewUserService(mockUserRepository)

	newUser := user.User{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane.doe@mail.com",
	}
	mockUserRepository.On("CreateUser", newUser).Return(newUser, nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = underTest.CreateUser(newUser)
	}
}

func BenchmarkUpdateUserById(b *testing.B) {
	mockUserRepository := new(mocks.Repository)
	underTest := user.NewUserService(mockUserRepository)

	updatedUser := users[0]
	updatedUser.FirstName = "Edward"
	mockUserRepository.On("UpdateUserById", id, mock.Anything).Return(updatedUser, nil)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = underTest.UpdateUserById(id, updatedUser)
	}
}

func BenchmarkDeleteUserById(b *testing.B) {
	mockUserRepository := new(mocks.Repository)
	underTest := user.NewUserService(mockUserRepository)

	mockUserRepository.On("DeleteUserById", id).Return(nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		underTest.DeleteUserById(id)
	}
}
