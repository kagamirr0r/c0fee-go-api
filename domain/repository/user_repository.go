package repository

import (
	"c0fee-api/domain/entity"

	"github.com/google/uuid"
)

type IUserRepository interface {
	GetById(user *entity.User, id uuid.UUID) error
	CreateUser(user *entity.User) error
}