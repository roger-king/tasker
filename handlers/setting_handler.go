package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/roger-king/tasker/models"
	"github.com/roger-king/tasker/services"
	"github.com/roger-king/tasker/utils"
)

// ListPluginSetting -
func ListPluginSetting(s *services.SettingService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()

		// Default type of plugin
		filters := map[string]interface{}{
			"type": "plugin",
		}

		if len(queries) > 0 {
			if len(queries["active"]) > 0 {
				boolean, _ := strconv.ParseBool(queries["active"][0])
				filters["active"] = boolean
			}

			if len(queries["skip"]) > 0 {
				i, _ := strconv.ParseInt(queries["skip"][0], 10, 64)
				filters["skip"] = i
			}
		}

		settings, err := s.ListPluginSettings(filters)

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
