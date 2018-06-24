package word

import (
	"net/http"
	"encoding/json"
	"ielts_study_backend/controllers/api"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	res := api.Response{"Welcome to my very first api", 200, nil}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}