package controllers

import (
	"encoding/json"
	"moodly/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRandomFeelingByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	f := models.GetRandomFeeling(category)
	res, _ := json.Marshal(f)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
