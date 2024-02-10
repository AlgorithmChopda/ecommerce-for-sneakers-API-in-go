package apperrors

import "net/http"

func MapError(err error) (statusCode int, errResponse error) {
	switch err.(type) {
	case UserAlreadyPresent:
		return http.StatusConflict, err
	case ProductNotFound:
		return http.StatusNotFound, err
	case CartAlreadyPresent:
		return http.StatusConflict, err
	default:
		return http.StatusInternalServerError, err
	}
}
