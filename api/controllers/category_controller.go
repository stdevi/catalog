package controllers

import (
	"catalog/api/models"
	"encoding/json"
	"net/http"
)

func (s *Server) GetCategories(w http.ResponseWriter, _ *http.Request) {
	cs, err := (&models.Category{}).FindAll(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
