package store

import (
	"gotraining/entities"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

// Repository Interface
type Repository interface {
	Create(store entities.Store) (entities.Store, error)
	Read(id int64) (entities.Store, error)
	List() ([]entities.Store, error)
	Update(id int64, store entities.Store) (entities.Store, error)
	Delete(id int64) error
}

// SetRepository function
func SetRepository(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (p *repo) Create(store entities.Store) (entities.Store, error) {
	err := p.db.Create(&store).Error
	return store, err
}

func (p *repo) Read(id int64) (entities.Store, error) {
	store := entities.Store{}
	err := p.db.Where("id = ?", id).First(&store).Error

	return store, err
}

func (p *repo) List() ([]entities.Store, error) {
	stores := []entities.Store{}
	err := p.db.Find(&stores).Error

	return stores, err
}

func (p *repo) Update(id int64, store entities.Store) (entities.Store, error) {
	err := p.db.Where("id = ?", id).First(&store).Updates(&store).Error

	return store, err
}

func (p *repo) Delete(id int64) error {
	var store entities.Store
	err := p.db.Where("id = ?", id).First(&store).Delete(&store).Error

	return err
}
