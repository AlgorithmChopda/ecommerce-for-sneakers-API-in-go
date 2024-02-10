package apperrors

type UserAlreadyPresent struct{}

func (u UserAlreadyPresent) Error() string {
	return "email already exists"
}

type ProductNotFound struct{}

func (p ProductNotFound) Error() string {
	return "no such product found"
}

type CartAlreadyPresent struct{}

func (p CartAlreadyPresent) Error() string {
	return "cart already present"
}
