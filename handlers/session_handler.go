package handlers

import (
	"errors"
	"net/http"

	"github.com/roger-king/tasker/utils"
)

func CheckSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode, _ := checkCookie(w, r)

		if statusCode != 200 {
			respondWithError(w, statusCode, utils.ProcessingError, errors.New("Failed to validate"))
			return
		}

		respondWithJSON(w, http.StatusOK, true)
		return
	}
}

func GetCurrentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getCurrentUser(r.Context())

		respondWithJSON(w, http.StatusOK, user)
		return
	}
}
