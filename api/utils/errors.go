package utils

import (
	"errors"
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

func FormatCategoryError(error error) string {
	if strings.Contains(error.Error(), "name") {
		return ErrInvalidCategoryName.Error()
	}

	return error.Error()
}

func FormatProductError(error error) string {
	if strings.Contains(error.Error(), "name") {
		return ErrInvalidProductName.Error()
	}
	if strings.Contains(error.Error(), "price") {
		return ErrInvalidProductPrice.Error()
	}

	return error.Error()
}
