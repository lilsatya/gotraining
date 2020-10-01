package product

import (
	"gotraining/entities"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

// Repository Interface
type Repository interface {
	Create(product entities.Product) (entities.Product, error)
	Read(id int) (entities.Product, error)
	List() ([]entities.Product, error)
	// Update(product entities.Product) (entities.Product, error)
	// Delete(id int) (entities.Product, error)
}

// SetRepository function
func SetRepository(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (p *repo) Create(product entities.Product) (entities.Product, error) {
	res := p.db.Create(&product)
	return product, res.Error
}

func (p *repo) Read(id int) (entities.Product, error) {
	product := entities.Product{}
	err := p.db.Where("id = ?", id).Find(&product).Error

	return product, err
}

func (p *repo) List() ([]entities.Product, error) {
	products := []entities.Product{}
	res := p.db.Find(&products)

	return products, res.Error
}
