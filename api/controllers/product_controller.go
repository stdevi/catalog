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

func (s *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.FormatProductError(w, err)
		return
	}

	p := models.Product{}
	if err := json.Unmarshal(b, &p); err != nil {
		utils.FormatProductError(w, err)
		return
	}

	if err := p.Validate(); err != nil {
		utils.FormatProductError(w, err)
		return
	}

	pCreated, err := p.Save(s.DB)
	if err != nil {
		utils.FormatProductError(w, err)
		return
	}

	utils.WriteResponse(w, pCreated)
}

func (s *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.FormatProductError(w, err)
		return
	}

	p := models.Product{}
	if err := json.Unmarshal(b, &p); err != nil {
		utils.FormatProductError(w, err)
		return
	}

	if err := p.Validate(); err != nil {
		utils.FormatProductError(w, err)
		return
	}

	pUpdated, err := p.Update(s.DB, uint(id))
	if err != nil {
		utils.FormatProductError(w, err)
		return
	}

	utils.WriteResponse(w, pUpdated)
}

func (s *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := (&models.Product{}).Delete(s.DB, uint(id)); err != nil {
		utils.FormatProductError(w, err)
		return
	}
}

func (s *Server) GetProducts(w http.ResponseWriter, _ *http.Request) {
	ps, err := (&models.Product{}).FindAll(s.DB)
	if err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	utils.WriteResponse(w, ps)
}

func (s *Server) GetProductsByCategoryId(w http.ResponseWriter, r *http.Request) {
	cid, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ps, err := (&models.Product{}).FindAllByCategoryId(s.DB, uint(cid))
	if err != nil {
		utils.FormatCategoryError(w, err)
		return
	}

	utils.WriteResponse(w, ps)
}
