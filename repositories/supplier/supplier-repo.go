package supplier

import (
	"gotraining/entities"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

// Repository Interface
type Repository interface {
	Create(Supplier entities.Supplier) (entities.Supplier, error)
	Read(id int) (entities.Supplier, error)
	List() ([]entities.Supplier, error)
	// Update(Supplier entities.Supplier) (entities.Supplier, error)
	// Delete(id int) (entities.Supplier, error)
}

// SetRepository function
func SetRepository(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (p *repo) Create(Supplier entities.Supplier) (entities.Supplier, error) {
	res := p.db.Create(&Supplier)
	return Supplier, res.Error
}

func (p *repo) Read(id int) (entities.Supplier, error) {
	Supplier := entities.Supplier{}
	err := p.db.Where("id = ?", id).Find(&Supplier).Error

	return Supplier, err
}

func (p *repo) List() ([]entities.Supplier, error) {
	Suppliers := []entities.Supplier{}
	res := p.db.Find(&Suppliers)

	return Suppliers, res.Error
}
