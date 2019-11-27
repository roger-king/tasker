package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/roger-king/tasker/models"
	"github.com/roger-king/tasker/services"
	"github.com/roger-king/tasker/utils"
)

// ListPluginSetting -
func ListPluginSetting(s *services.SettingService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		settings, err := s.ListPluginSettings()

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "", err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, settings)
		return
	}
}

// CreatePluginSetting -
func CreatePluginSetting(s *services.SettingService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input models.PluginSetting
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&input); err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.RequestError, err.Error())
			return
		}

		setting, err := s.CreatePluginSetting(&input)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, setting)
		return
	}
}

func ToggleActiveRepository(s *services.SettingService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input models.ToggleActiveSetting
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&input); err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.RequestError, err.Error())
			return
		}

		err := s.ToggleActiveSettingPluginRepo(&input)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, utils.ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, true)
		return
	}
}
