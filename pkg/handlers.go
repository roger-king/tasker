package pkg

import (
	"encoding/json"
	"net/http"
)

// ListTasks -
func ListTasks(t *TaskService) http.HandlerFunc {
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
func CreateTask(t *TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input NewInputTask
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&input); err != nil {
			respondWithError(w, http.StatusInternalServerError, RequestError, err.Error())
			return
		}

		tasks, err := t.Create(&input)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, tasks)
		return
	}
}

// // FindOneTask -
// func FindTask(t *TaskService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		task, err := t.Find(vars["taskID"])

// 		if err != nil {
// 			respondWithError(w, http.StatusInternalServerError, ProcessingError, err.Error())
// 			return
// 		}

// 		respondWithJSON(w, http.StatusOK, task)
// 		return
// 	}
// }

// // DisableTask -
// func DisableTask(session db.Database) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		defer session.Close()

// 		taskService := NewTaskService(session)
// 		err := taskService.Disable(vars["taskID"])

// 		if err != nil {
// 			respondWithError(w, http.StatusInternalServerError, ProcessingError, err.Error())
// 			return
// 		}

// 		respondWithJSON(w, http.StatusOK, "")
// 		return
// 	}
// }

// DeleteTask -
// func DeleteTask(t *TaskService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		ok, err := t.Delete(vars["taskID"])

// 		if err != nil {
// 			respondWithError(w, http.StatusInternalServerError, ProcessingError, err.Error())
// 			return
// 		}

// 		respondWithJSON(w, http.StatusOK, ok)
// 		return
// 	}
// }

// Response Helpers:
type errorHelper struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

type dataHelper struct {
	Data interface{} `json:"data"`
}

func respondWithError(w http.ResponseWriter, code int, errorType HTTPError, message string) {
	response, _ := json.Marshal(&errorHelper{Error: errorType.String(), Description: message})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(&dataHelper{Data: payload})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
