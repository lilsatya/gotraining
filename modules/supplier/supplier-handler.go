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
	supplier, err = h.repo.Read(int64(id))
	return
}

func (h *handler) list() (supplierList []entities.Supplier, err error) {
	supplierList, err = h.repo.List()
	return
}

func (h *handler) update(id int, data entities.Supplier) (supplier entities.Supplier, err error) {
	supplier, err = h.repo.Update(int64(id), data)
	return
}

func (h *handler) delete(id int) (err error) {
	err = h.repo.Delete(int64(id))
	return
}
