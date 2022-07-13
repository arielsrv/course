package infrastructure

import (
	"course/src/users/application"
	"course/src/users/domain"
	"encoding/json"
	"net/http"
)

type UserController struct {
	userService *application.UserService
}

func NewUserController(userService *application.UserService) *UserController {
	return &UserController{userService: userService}
}

func GetUser(writer http.ResponseWriter, _ *http.Request) {
	user := domain.User{Id: 1, Name: "Steve Jobs", Email: "stevejobs@apple.com"}
	json.NewEncoder(writer).Encode(user)
}
