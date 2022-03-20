package users

import (
	"course/src/users/application"
	"course/src/users/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) GetUser(int) domain.User {
	args := mock.Called()
	result := args.Get(0)
	return result.(domain.User)
}

func Test_Get_User(t *testing.T) {
	userRepository := new(MockUserRepository)
	userRepository.On("GetUser").Return(getUser())

	userService := application.NewUserService(userRepository)
	actual := userService.GetUser(1)

	userRepository.AssertExpectations(t)

	assert.Equal(t, 1, actual.Id)
	assert.Equal(t, "Name", actual.Name)
	assert.Equal(t, "name@email.com", actual.Email)
}

func getUser() domain.User {
	mockUser := domain.User{Id: 1, Name: "Name", Email: "name@email.com"}
	return mockUser
}
