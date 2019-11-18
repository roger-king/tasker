package handlers

import (
	"errors"
	"net/http"

	"github.com/roger-king/tasker/utils"
)

func CheckSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode := checkCookie(w, r)

		if statusCode != 200 {
			respondWithError(w, statusCode, utils.ProcessingError, errors.New("Failed to validate"))
			return
		}

		respondWithJSON(w, http.StatusOK, true)
		return
	}
}
