package user

import "event/entities"

type User interface {
	InsertUser(newUser entities.User) (entities.User, error)
	GetUserID(ID int) (entities.User, error)
	UpdateUser(ID int, update entities.User) (entities.User, error)
	DeleteUser(ID int) (entities.User, error)
	Login(email string, password string) (entities.User, error)
}
