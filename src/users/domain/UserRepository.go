package domain

type UserRepository interface {
	GetUser(id int64) *User
}
