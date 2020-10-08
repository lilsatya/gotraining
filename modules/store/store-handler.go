package store

import (
	"gotraining/config"
	"gotraining/entities"
	"gotraining/repositories/store"
)

type (
	handler struct {
		repo store.Repository
	}
)

var instance *handler

func newHandler() *handler {
	if instance == nil {
		db := config.OpenDB()
		instance = new(handler)
		instance.repo = store.SetRepository(db)
	}

	return instance
}

func (h *handler) create(data entities.Store) (store entities.Store, err error) {
	store, err = h.repo.Create(data)
	return
}

func (h *handler) read(id int) (store entities.Store, err error) {
	store, err = h.repo.Read(int64(id))
	return
}

func (h *handler) list() (storeList []entities.Store, err error) {
	storeList, err = h.repo.List()
	return
}

func (h *handler) update(id int, data entities.Store) (store entities.Store, err error) {
	store, err = h.repo.Update(int64(id), data)
	return
}

func (h *handler) delete(id int) (err error) {
	err = h.repo.Delete(int64(id))
	return
}
