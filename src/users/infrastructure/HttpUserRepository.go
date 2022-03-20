package infrastructure

import (
	"course/src/users/domain"
)

type HttpUserRepository struct {
	user *domain.User
}

func (httpUserRepository HttpUserRepository) GetUser(int64) *domain.User {
	httpUserRepository.user = new(domain.User)

	httpUserRepository.user.Id = 1
	httpUserRepository.user.Name = "Steve Jobs"
	httpUserRepository.user.Email = "stevejobs@apple.com"

	return httpUserRepository.user
}
