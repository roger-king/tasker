package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roger-king/tasker/models"
	"github.com/roger-king/tasker/services"
	"github.com/roger-king/tasker/utils"
)

// ListTasks -
func ListTasks(t *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := t.List()

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "", err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, tasks)
		return
	}
}

// CreateTask -
func CreateTask(t *services.TaskService, u *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input models.NewInputTask
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&input); err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.RequestError, err.Error())
			return
		}

		currUser := r.Context().Value(ContextKey("user")).(models.User)
		foundUser, err := u.FindUser(currUser.UserName)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "", err)
			return
		}

		token, err := foundUser.GetAccessToken()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "", err)
			return
		}

		gAPI := services.NewGithubAPIService(token)
		gAPI.DownloadTaggedAssets()

		tasks, err := t.Create(&input)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, tasks)
		return
	}
}

// FindOneTask -
func FindTask(t *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskId := vars["taskID"]

		if len(taskId) == 0 {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, errors.New("no task id was supplied").Error())
			return
		}

		task, err := t.Find(taskId)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, task)
		return
	}
}

// DisableTask -
func DisableTask(t *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskId := vars["taskID"]

		if len(taskId) == 0 {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, errors.New("no task id was supplied").Error())
			return
		}

		err := t.Disable(taskId)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, true)
		return
	}
}

// DeleteTask -
func DeleteTask(t *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskId := vars["taskID"]

		if len(taskId) == 0 {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, errors.New("no task id was supplied").Error())
			return
		}

		err := t.Delete(taskId)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, true)
		return
	}
}
