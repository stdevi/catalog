package modeltests

import (
	"catalog/api/models"
	"catalog/api/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFindAllCategories(t *testing.T) {
	if err := refreshCategoryTable(); err != nil {
		log.Fatal(err)
	}

	if err := seedCategories(); err != nil {
		log.Fatal(err)
	}

	cs, err := category.FindAll(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(cs), 2)
}

func TestFindCategoryById(t *testing.T) {
	if err := refreshCategoryTable(); err != nil {
		log.Fatal(err)
	}

	seedC, err := seedSingleCategory()
	if err != nil {
		log.Fatal(err)
	}

	c, err := category.FindById(server.DB, 1)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, c.ID, seedC.ID)
	assert.Equal(t, c.Name, seedC.Name)
}

func TestCategoryValidation(t *testing.T) {
	c := models.Category{
		Name: "",
	}

	err := c.Validate()

	assert.EqualError(t, err, utils.ErrInvalidCategoryName.Error())
}

func TestSaveCategory(t *testing.T) {
	if err := refreshCategoryTable(); err != nil {
		log.Fatal(err)
	}

	newC := models.Category{
		Name: "Garden",
	}

	savedC, err := newC.Save(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, newC.Name, savedC.Name)
}

func TestUpdateCategory(t *testing.T) {
	if err := refreshCategoryTable(); err != nil {
		log.Fatal(err)
	}

	seedC, err := seedSingleCategory()
	if err != nil {
		log.Fatal(err)
	}

	newC := models.Category{
		Name: "Garden",
	}

	updatedC, err := newC.Update(server.DB, seedC.ID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, newC.Name, updatedC.Name)
}

func TestDeleteCategory(t *testing.T) {
	if err := refreshCategoryTable(); err != nil {
		log.Fatal(err)
	}

	seedC, err := seedSingleCategory()
	if err != nil {
		log.Fatal(err)
	}

	err = category.Delete(server.DB, seedC.ID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, nil, err)
}
