package controllers

import (
	"catalog/api/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
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

	WriteResponse(w, cCreated)
}

func (s *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	cUpdated, err := c.Update(s.DB, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, cUpdated)
}

func (s *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := (&models.Category{}).Delete(s.DB, uint(id)); err != nil {
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

	WriteResponse(w, cs)
}
