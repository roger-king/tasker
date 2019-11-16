package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roger-king/tasker/services"
	"github.com/roger-king/tasker/utils"
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
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		if len(resp.Error) > 0 {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, resp)
			return
		}

		respondWithJSON(w, http.StatusOK, resp)
		return
	}
}

func FetchUserClientIDHandler(gh *services.GithubService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := gh.FetchClientID(utils.GithubScopeType("user"))

		if len(id.ClientID) <= 0 {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, errors.New("cannot find github client id").Error())
			return
		}

		respondWithJSON(w, http.StatusOK, id)
		return
	}
}
