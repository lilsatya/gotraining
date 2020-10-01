package product

import (
	"gotraining/config"
	"gotraining/entities"
	"gotraining/repositories/product"
)

type (
	handler struct {
		repo product.Repository
	}
)

var instance *handler

func newHandler() *handler {
	if instance == nil {
		db := config.OpenDB()
		instance = new(handler)
		instance.repo = product.SetRepository(db)
	}

	return instance
}

func (h *handler) create(data entities.Product) (product entities.Product, err error) {
	product, err = h.repo.Create(data)
	return
}

func (h *handler) read(id int) (product entities.Product, err error) {
	product, err = h.repo.Read(id)
	return
}

func (h *handler) list() (productList []entities.Product, err error) {
	productList, err = h.repo.List()
	return
}
