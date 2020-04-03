package controllers

import (
	"catalog/api/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	c := models.Category{}
	if err := json.Unmarshal(b, &c); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	cCreated, err := c.Save(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cCreated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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
