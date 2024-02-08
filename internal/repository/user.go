package repository

import (
	"time"
)

type User struct {
	ID          int       `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	DateOfBirth time.Time `db:"date_of_birth"`
	MobileNo    int       `db:"mobile_no"`
	Address     string    `db:"address"`
	City        string    `db:"city"`
	PostalCode  int       `db:"postal_code"`
	CompanyID   int       `db:"company_id"`
	RoleID      int       `db:"role_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type UserRepository interface {
	CreateUser(userInfo []any) error
	CheckUserWithEmailAndPassword(email, password string) error
	GetIdRoleAndPassword(email string) (int, int, string, error)
}
