package domain

type UserRepository interface {
	GetUser(id int) User
}
