package supplier

import (
	"gotraining/config"
	"gotraining/entities"
	"gotraining/repositories/supplier"
)

type (
	handler struct {
		repo supplier.Repository
	}
)

var instance *handler

func newHandler() *handler {
	if instance == nil {
		db := config.OpenDB()
		instance = new(handler)
		instance.repo = supplier.SetRepository(db)
	}

	return instance
}

func (h *handler) create(data entities.Supplier) (supplier entities.Supplier, err error) {
	supplier, err = h.repo.Create(data)
	return
}

func (h *handler) read(id int) (supplier entities.Supplier, err error) {
	supplier, err = h.repo.Read(id)
	return
}

func (h *handler) list() (supplierList []entities.Supplier, err error) {
	supplierList, err = h.repo.List()
	return
}
