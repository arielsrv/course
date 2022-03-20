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

func (userService *UserService) GetUser(id int64) *UserDto {
	user := userService.userRepository.GetUser(id)

	userDto := new(UserDto)
	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.Email = user.Email

	return userDto
}
