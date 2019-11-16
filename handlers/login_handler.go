package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roger-king/tasker/utils"
)

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		code := vars["code"]

		if len(code) == 0 {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, errors.New("no code was supplied").Error())
			return
		}

		respondWithJSON(w, http.StatusOK, true)
		return
	}
}
