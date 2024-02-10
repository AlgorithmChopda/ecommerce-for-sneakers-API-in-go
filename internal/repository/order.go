package repository

type OrderRepository interface {
	Create(userId int) (int, error)
	IsOrderPresent(userId int) (bool, error)
}
