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
	Create(supplier entities.Supplier) (entities.Supplier, error)
	Read(id int64) (entities.Supplier, error)
	List() ([]entities.Supplier, error)
	Update(id int64, supplier entities.Supplier) (entities.Supplier, error)
	Delete(id int64) error
}

// SetRepository function
func SetRepository(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (p *repo) Create(supplier entities.Supplier) (entities.Supplier, error) {
	err := p.db.Create(&supplier).Error
	return supplier, err
}

func (p *repo) Read(id int64) (entities.Supplier, error) {
	supplier := entities.Supplier{}
	err := p.db.Where("id = ?", id).First(&supplier).Error

	return supplier, err
}

func (p *repo) List() ([]entities.Supplier, error) {
	suppliers := []entities.Supplier{}
	err := p.db.Find(&suppliers).Error

	return suppliers, err
}

func (p *repo) Update(id int64, supplier entities.Supplier) (entities.Supplier, error) {
	err := p.db.Where("id = ?", id).First(&supplier).Updates(&supplier).Error

	return supplier, err
}

func (p *repo) Delete(id int64) error {
	var supplier entities.Supplier
	err := p.db.Where("id = ?", id).First(&supplier).Delete(&supplier).Error

	return err
}
