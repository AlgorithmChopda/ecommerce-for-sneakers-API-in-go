package apperrors

type UserAlreadyPresent struct{}

func (u UserAlreadyPresent) Error() string {
	return "email already exists"
}

type ProductNotFound struct{}

func (p ProductNotFound) Error() string {
	return "no such product found"
}
