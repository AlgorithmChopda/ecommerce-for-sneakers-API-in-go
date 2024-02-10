package helpers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPathParameter(r *http.Request, parameter string) (int, error) {
	pathId := mux.Vars(r)[parameter]

	id, err := strconv.Atoi(pathId)
	if err != nil {
		return -1, errors.New("Invalid parameter in path")
	}

	return id, nil
}
