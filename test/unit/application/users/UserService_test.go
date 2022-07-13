package users

import (
	"course/src/users/application"
	"course/src/users/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UnitTestSuite struct {
	suite.Suite
	userRepository *MockUserRepository
}

func (suite *UnitTestSuite) SetupTest() {
	suite.userRepository = new(MockUserRepository)
}

func TestUnitUserService(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) GetUser(int) domain.User {
	args := mock.Called()
	result := args.Get(0)
	return result.(domain.User)
}

func (suite *UnitTestSuite) Test_Get_User() {
	suite.userRepository.On("GetUser").Return(getUser())

	userService := application.NewUserService(suite.userRepository)
	actual := userService.GetUser(1)

	suite.userRepository.AssertExpectations(suite.T())

	suite.Equal(1, actual.Id)
	suite.Equal("Name", actual.Name)
	suite.Equal("name@email.com", actual.Email)
}

func getUser() domain.User {
	mockUser := domain.User{Id: 1, Name: "Name", Email: "name@email.com"}
	return mockUser
}
