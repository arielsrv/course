package main

import (
	"course/src/users/application"
	"course/src/users/domain"
	"course/src/users/infrastructure"
	"fmt"
)

func main() {
	var userRepository domain.UserRepository
	userRepository = new(infrastructure.HttpUserRepository)

	var userService = application.NewUserService(userRepository)
	var userDto = userService.GetUser(1)

	fmt.Println(userDto.Id)
	fmt.Println(userDto.Name)
	fmt.Println(userDto.Email)
}
