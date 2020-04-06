package controllertests

import (
	"bytes"
	"catalog/api/models"
	"catalog/api/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {
	if err := refreshProductAndCategoryTables(); err != nil {
		log.Fatal(err)
	}

	if err := seedCategoryAndProducts(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/products", nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetProducts)
	handler.ServeHTTP(rr, req)

	var ps []models.Product
	if err := json.Unmarshal([]byte(rr.Body.String()), &ps); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, len(ps), 3)
}

func TestGetProductsByCategory(t *testing.T) {
	if err := refreshProductAndCategoryTables(); err != nil {
		log.Fatal(err)
	}

	if err := seedCategoryAndProducts(); err != nil {
		log.Fatal(err)
	}

	cases := []struct {
		id   string
		code int
	}{
		{id: "1", code: http.StatusOK},
		{id: "2", code: http.StatusOK},
		{id: "42", code: http.StatusNotFound},
		{id: "my_id", code: http.StatusBadRequest},
	}

	for _, input := range cases {
		req, err := http.NewRequest("GET", "/api/products/category/", nil)
		if err != nil {
			log.Fatal(err)
		}

		req = mux.SetURLVars(req, map[string]string{"id": input.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.GetProductsByCategoryId)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, input.code, rr.Code)
		if input.code == http.StatusOK {
			var ps []models.Product
			if err := json.Unmarshal([]byte(rr.Body.String()), &ps); err != nil {
				log.Fatal(err)
			}

			if input.id == "1" {
				assert.Equal(t, len(ps), 2)
			}
			if input.id == "2" {
				assert.Equal(t, len(ps), 1)
			}
		}
		if input.id == "42" {
			assert.Error(t, utils.ErrCategoryNotFound)
		}
	}
}

func TestCreateProduct(t *testing.T) {
	if err := refreshProductAndCategoryTables(); err != nil {
		log.Fatal(err)
	}

	seedC, err := seedSingleCategory()
	if err != nil {
		log.Fatal(err)
	}

	cases := []struct {
		body        string
		name        string
		description string
		price       float32
		categoryId  uint
		code        int
	}{
		{
			body:        `{"name": "test", "description": "desc", "price": 42, "category_id": 1}`,
			name:        "test",
			description: "desc",
			price:       42,
			categoryId:  seedC.ID,
			code:        http.StatusOK,
		},
		{
			body:        `{"name": "", "description": "test desc", "price": 42, "category_id": 1}`,
			name:        "",
			description: "test desc",
			price:       42,
			categoryId:  seedC.ID,
			code:        http.StatusUnprocessableEntity,
		},
		{
			body:        `{"name": "test name", "description": "test desc", "price": -300, "category_id": 1}`,
			name:        "test name",
			description: "test desc",
			price:       -300,
			categoryId:  seedC.ID,
			code:        http.StatusUnprocessableEntity,
		},
		{
			body:        `{"name": "test name", "description": "test desc", "price": 42, "category_id": 42}`,
			name:        "test name",
			description: "test desc",
			price:       42,
			categoryId:  uint(42),
			code:        http.StatusNotFound,
		},
	}

	for _, input := range cases {
		req, err := http.NewRequest("POST", "/api/products", bytes.NewBufferString(input.body))
		if err != nil {
			log.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateProduct)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, input.code)
		if input.code == http.StatusOK {
			var p models.Product
			if err := json.Unmarshal([]byte(rr.Body.String()), &p); err != nil {
				log.Fatal(err)
			}
			assert.Equal(t, input.name, p.Name)
			assert.Equal(t, input.description, p.Description)
			assert.Equal(t, input.price, p.Price)
			assert.Equal(t, input.categoryId, p.CategoryID)
		}
		if input.code == http.StatusUnprocessableEntity {
			if input.name == "" {
				assert.Error(t, utils.ErrInvalidProductName)
			}
			if input.price < 0 {
				assert.Error(t, utils.ErrInvalidProductPrice)
			}
		}
		if input.categoryId != seedC.ID {
			assert.Error(t, utils.ErrCategoryNotFound)
		}
	}
}

func TestUpdateProduct(t *testing.T) {
	if err := refreshProductAndCategoryTables(); err != nil {
		log.Fatal(err)
	}

	p, err := seedSingleCategoryAndProduct()
	if err != nil {
		log.Fatal(err)
	}

	cases := []struct {
		body        string
		id          string
		name        string
		description string
		price       float32
		categoryId  uint
		code        int
	}{
		{
			body:        `{"name": "test", "description": "desc", "price": 42, "category_id": 1}`,
			id:          fmt.Sprint(p.ID),
			name:        "test",
			description: "desc",
			price:       42,
			categoryId:  p.CategoryID,
			code:        http.StatusOK,
		},
		{
			body:        `{"name": "", "description": "test desc", "price": 42, "category_id": 1}`,
			id:          fmt.Sprint(p.ID),
			name:        "",
			description: "test desc",
			price:       42,
			categoryId:  p.CategoryID,
			code:        http.StatusUnprocessableEntity,
		},
		{
			body:        `{"name": "test name", "description": "test desc", "price": -300, "category_id": 1}`,
			id:          fmt.Sprint(p.ID),
			name:        "test name",
			description: "test desc",
			price:       -300,
			categoryId:  p.CategoryID,
			code:        http.StatusUnprocessableEntity,
		},
		{
			body:        `{"name": "test name", "description": "test desc", "price": 42, "category_id": 42}`,
			id:          fmt.Sprint(p.ID),
			name:        "test name",
			description: "test desc",
			price:       42,
			categoryId:  uint(42),
			code:        http.StatusNotFound,
		},
		{
			body:        `{"name": "test name", "description": "test desc", "price": 42, "category_id": 42}`,
			id:          "42",
			name:        "test name",
			description: "test desc",
			price:       42,
			categoryId:  p.CategoryID,
			code:        http.StatusNotFound,
		},
	}

	for _, input := range cases {
		req, err := http.NewRequest("PUT", "/api/products", bytes.NewBufferString(input.body))
		if err != nil {
			log.Fatal(err)
		}

		req = mux.SetURLVars(req, map[string]string{"id": input.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.UpdateProduct)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, input.code)
		if input.code == http.StatusOK {
			var p models.Product
			if err := json.Unmarshal([]byte(rr.Body.String()), &p); err != nil {
				log.Fatal(err)
			}
			assert.Equal(t, input.name, p.Name)
			assert.Equal(t, input.description, p.Description)
			assert.Equal(t, input.price, p.Price)
			assert.Equal(t, input.categoryId, p.CategoryID)
		}
		if input.code == http.StatusUnprocessableEntity {
			if input.name == "" {
				assert.Error(t, utils.ErrInvalidProductName)
			}
			if input.price < 0 {
				assert.Error(t, utils.ErrInvalidProductPrice)
			}
		}
		if input.categoryId == 42 {
			assert.Error(t, utils.ErrCategoryNotFound)
		}
	}
}

func TestDeleteProduct(t *testing.T) {
	if err := refreshProductAndCategoryTables(); err != nil {
		return
	}

	p, err := seedSingleCategoryAndProduct()
	if err != nil {
		return
	}

	cases := []struct {
		id   string
		code int
	}{
		{id: fmt.Sprint(p.ID), code: http.StatusOK},
		{id: "42", code: http.StatusNotFound},
	}
	for _, input := range cases {
		req, err := http.NewRequest("DELETE", "/api/products/", nil)
		if err != nil {
			return
		}

		req = mux.SetURLVars(req, map[string]string{"id": input.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.DeleteProduct)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, input.code, rr.Code)
		if input.code == http.StatusNotFound {
			assert.Error(t, utils.ErrProductNotFound)
		}
	}
}
