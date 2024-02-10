package helpers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPathParameterId(r *http.Request) (int, error) {
	pathId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(pathId)
	if err != nil {
		return -1, errors.New("Invalid Id in path")
	}

	return id, nil
}
