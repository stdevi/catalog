package controllers

import (
	"catalog/api/models"
	"catalog/api/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	c := models.Category{}
	if err := json.Unmarshal(b, &c); err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	if err := c.Validate(); err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	cCreated, err := c.Save(s.DB)
	if err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	utils.WriteResponse(w, cCreated)
}

func (s *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	c := models.Category{}
	if err := json.Unmarshal(b, &c); err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	if err := c.Validate(); err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	cUpdated, err := c.Update(s.DB, uint(id))
	if err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	utils.WriteResponse(w, cUpdated)
}

func (s *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := (&models.Category{}).Delete(s.DB, uint(id)); err != nil {
		utils.FormatCategoryError(w, err)
		return
	}
}

func (s *Server) GetCategories(w http.ResponseWriter, _ *http.Request) {
	cs, err := (&models.Category{}).FindAll(s.DB)
	if err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	utils.WriteResponse(w, cs)
}
