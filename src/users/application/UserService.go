package application

import (
	"course/src/users/domain"
)

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) GetUser(id int) UserDto {
	user := service.userRepository.GetUser(id)
	userDto := UserDto{Id: user.Id, Name: user.Name, Email: user.Email}
	return userDto
}
