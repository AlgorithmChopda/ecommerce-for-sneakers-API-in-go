package apperrors

type UserAlreadyPresent struct{}

func (u UserAlreadyPresent) Error() string {
	return "Email already present"
}
