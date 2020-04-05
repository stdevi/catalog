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

func TestGetCategories(t *testing.T) {
	if err := refreshCategoryTable(); err != nil {
		log.Fatal(err)
	}

	if err := seedCategories(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/categories", nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetCategories)
	handler.ServeHTTP(rr, req)

	var cs []models.Category
	if err := json.Unmarshal([]byte(rr.Body.String()), &cs); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, len(cs), 2)
}

func TestCreateCategory(t *testing.T) {
	if err := refreshCategoryTable(); err != nil {
		log.Fatal(err)
	}

	cases := []struct {
		body string
		name string
		code int
	}{
		{
			body: `{"name": "test category"}`,
			name: "test category",
			code: 200,
		},
		{
			body: `{"name": ""}`,
			name: "",
			code: 500,
		},
	}

	for _, input := range cases {
		req, err := http.NewRequest("POST", "/api/categories", bytes.NewBufferString(input.body))
		if err != nil {
			log.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateCategory)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, input.code)
		if input.code == 200 {
			var c models.Category
			if err := json.Unmarshal([]byte(rr.Body.String()), &c); err != nil {
				log.Fatal(err)
			}
			assert.Equal(t, input.name, c.Name)
		}
		if input.code == 500 {
			assert.Error(t, utils.ErrInvalidCategoryName)
		}
	}
}

func TestUpdateCategory(t *testing.T) {
	if err := refreshCategoryTable(); err != nil {
		log.Fatal(err)
	}

	seedC, err := seedSingleCategory()
	if err != nil {
		log.Fatal(err)
	}

	cases := []struct {
		body string
		name string
		id   string
		code int
	}{
		{
			body: `{"name": "test category"}`,
			name: "test category",
			id:   fmt.Sprint(seedC.ID),
			code: 200,
		},
		{
			body: `{"name": "test category"}`,
			name: "test category",
			id:   "42",
			code: 500,
		},
		{
			body: `{"name": ""}`,
			name: "",
			id:   fmt.Sprint(seedC.ID),
			code: 500,
		},
		{
			body: `{"name": "test category"}`,
			name: "test category",
			id:   "my_id",
			code: 400,
		},
	}

	for _, input := range cases {
		req, err := http.NewRequest("PUT", "/api/categories/", bytes.NewBufferString(input.body))
		if err != nil {
			log.Fatal(err)
		}

		req = mux.SetURLVars(req, map[string]string{"id": input.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.UpdateCategory)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, input.code, rr.Code)
		if input.code == 200 {
			var c models.Category
			if err := json.Unmarshal([]byte(rr.Body.String()), &c); err != nil {
				log.Fatal(err)
			}
			assert.Equal(t, input.name, c.Name)
		}
		if input.code == 500 {
			if input.name == "" {
				assert.Error(t, utils.ErrInvalidCategoryName)
			}
			if input.id != fmt.Sprint(seedC.ID) {
				assert.Error(t, utils.ErrCategoryNotFound)
			}
		}
	}
}

func TestDeleteCategory(t *testing.T) {
	if err := refreshCategoryTable(); err != nil {
		log.Fatal(err)
	}

	seedC, err := seedSingleCategory()
	if err != nil {
		log.Fatal(err)
	}

	cases := []struct {
		id   string
		code int
	}{
		{
			id:   fmt.Sprint(seedC.ID),
			code: 200,
		},
		{
			id:   "42",
			code: 500,
		},
		{
			id:   "my_id",
			code: 400,
		},
	}

	for _, input := range cases {
		req, err := http.NewRequest("DELETE", "/api/categories/", nil)
		if err != nil {
			log.Fatal(err)
		}

		req = mux.SetURLVars(req, map[string]string{"id": input.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.DeleteCategory)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, input.code, rr.Code)
		if input.code == 500 {
			assert.Error(t, utils.ErrCategoryNotFound)
		}
	}
}
