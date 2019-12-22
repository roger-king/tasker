package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roger-king/tasker/models"
	"github.com/roger-king/tasker/services"
	"github.com/roger-king/tasker/utils"
)

// LoginHandler -
func LoginHandler(gh *services.GithubAuthService, us *services.UserService) http.HandlerFunc {
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

		// Fetching for github user
		api := services.NewGithubAPIService(resp.AccessToken)
		user, err := api.GetUser()

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		createdUser, err := us.CreateUser(&models.NewUserInput{
			Email:       user.GetEmail(),
			UserName:    user.GetLogin(),
			Name:        user.GetName(),
			AccessToken: resp.AccessToken,
			Bio:         user.GetBio(),
			GitHubURL:   user.GetHTMLURL(),
		})

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		token, expiresAt, err := services.GenerateJWTToken(createdUser)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "tasker-user",
			Value:    token,
			Expires:  expiresAt,
			HttpOnly: true,
			Path:     "/",
		})

		respondWithJSON(w, http.StatusOK, "OK")
		return
	}
}

func FetchUserClientIDHandler(gh *services.GithubAuthService) http.HandlerFunc {
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
