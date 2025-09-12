package user

import "github.com/google/uuid"

type IUserRepository interface {
	GetById(user *Entity, id uuid.UUID) error
	CreateUser(user *Entity) error
}
