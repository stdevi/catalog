package utils

import (
	"errors"
	"net/http"
	"strings"
)

var (
	// Category errors
	ErrCategoryNotFound       = errors.New("category not found")
	ErrInvalidCategoryName    = errors.New("invalid category name")
	ErrCategoryDeletionFailed = errors.New("category deletion failed")

	// Product errors
	ErrProductNotFound       = errors.New("product not found")
	ErrInvalidProductName    = errors.New("invalid product name")
	ErrInvalidProductPrice   = errors.New("invalid product price")
	ErrProductDeletionFailed = errors.New("product deletion failed")
)

func FormatCategoryError(w http.ResponseWriter, error error) {
	if strings.Contains(error.Error(), "name") {
		http.Error(w, ErrInvalidCategoryName.Error(), http.StatusUnprocessableEntity)
		return
	}
	if strings.Contains(error.Error(), "not found") {
		http.Error(w, ErrCategoryNotFound.Error(), http.StatusNotFound)
		return
	}

	http.Error(w, error.Error(), http.StatusInternalServerError)
}

func FormatProductError(w http.ResponseWriter, error error) {
	if strings.Contains(error.Error(), "name") {
		http.Error(w, ErrInvalidProductName.Error(), http.StatusUnprocessableEntity)
		return
	}
	if strings.Contains(error.Error(), "price") {
		http.Error(w, ErrInvalidProductPrice.Error(), http.StatusUnprocessableEntity)
		return
	}
	if strings.Contains(error.Error(), "category not found") {
		http.Error(w, ErrCategoryNotFound.Error(), http.StatusNotFound)
		return
	}
	if strings.Contains(error.Error(), "product not found") {
		http.Error(w, ErrProductNotFound.Error(), http.StatusNotFound)
		return
	}

	http.Error(w, error.Error(), http.StatusInternalServerError)
}
