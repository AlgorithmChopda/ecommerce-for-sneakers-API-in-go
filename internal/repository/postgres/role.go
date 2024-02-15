package repository

import (
	"database/sql"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type roleStore struct {
	DB *sql.DB
}

func NewRoleRepository(db *sql.DB) repository.RoleRepository {
	return &roleStore{
		DB: db,
	}
}

func (role *roleStore) GetRoleId(roleName string) (int, error) {
	q := "select id from role where name = $1"
	row := role.DB.QueryRow(q, roleName)

	var id int
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
