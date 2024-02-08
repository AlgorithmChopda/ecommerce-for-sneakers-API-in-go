package repository

type RoleRepository interface {
	GetRoleId(roleName string) (int, error)
}
