package modeltests

import (
	"catalog/api/models"
	"catalog/api/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFindAllProducts(t *testing.T) {
	if err := refreshProductAndCategoryTables(); err != nil {
		log.Fatal(err)
	}

	if err := seedCategoryAndProducts(); err != nil {
		log.Fatal(err)
	}

	ps, err := product.FindAll(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(ps), 3)
}

func TestFindProductByCategoryId(t *testing.T) {
	if err := refreshProductAndCategoryTables(); err != nil {
		log.Fatal(err)
	}

	if err := seedCategoryAndProducts(); err != nil {
		log.Fatal(err)
	}

	ps, err := product.FindAllByCategoryId(server.DB, 1)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(ps), 2)
}

func TestProductNameValidation(t *testing.T) {
	p := models.Product{
		Name: "",
	}

	err := p.Validate()

	assert.EqualError(t, err, utils.ErrInvalidProductName.Error())
}

func TestProductPriceValidation(t *testing.T) {
	p := models.Product{
		Name:  "Mouse",
		Price: -30,
	}

	err := p.Validate()

	assert.EqualError(t, err, utils.ErrInvalidProductPrice.Error())
}

func TestSaveProduct(t *testing.T) {
	if err := refreshProductAndCategoryTables(); err != nil {
		log.Fatal(err)
	}

	c, err := seedSingleCategory()
	if err != nil {
		log.Fatal(err)
	}

	newP := models.Product{
		Name:        "Mouse",
		Description: "Best mouse ever",
		Price:       30.98,
		CategoryID:  c.ID,
	}

	savedP, err := newP.Save(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, newP.Name, savedP.Name)
	assert.Equal(t, newP.Description, savedP.Description)
	assert.Equal(t, newP.Price, savedP.Price)
	assert.Equal(t, newP.CategoryID, savedP.CategoryID)
}

func TestUpdateProduct(t *testing.T) {
	if err := refreshProductAndCategoryTables(); err != nil {
		log.Fatal(err)
	}

	seedP, err := seedSingleCategoryAndProduct()
	if err != nil {
		log.Fatal(err)
	}

	newP := models.Product{
		Name:        "Updated Product",
		Description: "Updated Product info",
		Price:       422,
		CategoryID:  seedP.CategoryID,
	}

	updatedP, err := newP.Update(server.DB, seedP.ID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, newP.Name, updatedP.Name)
	assert.Equal(t, newP.Description, updatedP.Description)
	assert.Equal(t, newP.Price, updatedP.Price)
	assert.Equal(t, newP.CategoryID, updatedP.CategoryID)
}

func TestDeleteProduct(t *testing.T) {
	if err := refreshProductAndCategoryTables(); err != nil {
		log.Fatal(err)
	}

	seedP, err := seedSingleCategoryAndProduct()
	if err != nil {
		log.Fatal(err)
	}

	err = product.Delete(server.DB, seedP.ID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, nil, err)
}
