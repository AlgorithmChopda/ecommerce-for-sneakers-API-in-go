package apperrors

type UserAlreadyPresent struct{}

func (u UserAlreadyPresent) Error() string {
	return "user already exists"
}

type NotFoundError struct{ Message string }

func (p NotFoundError) Error() string {
	return p.Message
}

type CartAlreadyPresent struct{}

func (p CartAlreadyPresent) Error() string {
	return "cart already present"
}

type InsufficientProductQuantity struct{}

func (i InsufficientProductQuantity) Error() string {
	return "insufficient product quantity"
}

type UnauthorizedAccess struct{ Message string }

func (u UnauthorizedAccess) Error() string {
	return u.Message
}

type EmptyError struct{ Message string }

func (e EmptyError) Error() string {
	return e.Message
}
