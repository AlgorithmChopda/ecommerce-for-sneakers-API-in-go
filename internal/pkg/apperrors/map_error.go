package apperrors

import "net/http"

func MapError(err error) (statusCode int, errResponse error) {
	switch err.(type) {
	case UserAlreadyPresent:
		return http.StatusConflict, err
	case NotFoundError:
		return http.StatusNotFound, err
	case CartAlreadyPresent:
		return http.StatusConflict, err
	case InsufficientProductQuantity:
		return http.StatusBadRequest, err
	case UnauthorizedAccess:
		return http.StatusUnauthorized, err
	case EmptyError:
		return http.StatusBadRequest, err
	default:
		return http.StatusInternalServerError, err
	}
}
