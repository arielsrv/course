package infrastructure

import (
	"course/src/users/domain"
)

type HttpUserRepository struct {
	user domain.User
}

func (repository HttpUserRepository) GetUser(int64) domain.User {
	user := domain.User{Id: 1, Name: "Steve Jobs", Email: "stevejobs@apple.com"}
	return user
}
