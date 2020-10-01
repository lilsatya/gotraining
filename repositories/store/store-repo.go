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
	Create(Store entities.Store) (entities.Store, error)
	Read(id int) (entities.Store, error)
	List() ([]entities.Store, error)
	// Update(Store entities.Store) (entities.Store, error)
	// Delete(id int) (entities.Store, error)
}

// SetRepository function
func SetRepository(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (p *repo) Create(Store entities.Store) (entities.Store, error) {
	res := p.db.Create(&Store)
	return Store, res.Error
}

func (p *repo) Read(id int) (entities.Store, error) {
	Store := entities.Store{}
	err := p.db.Where("id = ?", id).Find(&Store).Error

	return Store, err
}

func (p *repo) List() ([]entities.Store, error) {
	Stores := []entities.Store{}
	res := p.db.Find(&Stores)

	return Stores, res.Error
}
