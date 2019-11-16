package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roger-king/tasker/services"
	"github.com/roger-king/tasker/utils"
	"github.com/sirupsen/logrus"
)

func LoginHandler(gh *services.GithubService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		code := vars["code"]

		if len(code) == 0 {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, errors.New("no code was supplied").Error())
			return
		}

		resp, err := gh.GetAccessToken(code)

		if err != nil {
			logrus.Info(err)
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, resp)
		return
	}
}
