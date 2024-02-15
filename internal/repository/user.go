package repository

import (
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
)

type UserRepository interface {
	CreateUser(userInfo []any) error
	CheckUserWithEmailAndPassword(email, password string) error
	GetIdRoleAndPassword(email string) (int, int, string, error)
	IsUserWithEmailPresent(email string) bool
	GetUserList(roleId int) ([]dto.UserResponseObject, error)
	GetUserProfile(userId int) (dto.UserResponseObject, error)
}
