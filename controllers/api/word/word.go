package word

import (
	"net/http"
	"encoding/json"
	"ielts_study_backend/controllers"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	res := controllers.Response{"hello there", 200, nil}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}