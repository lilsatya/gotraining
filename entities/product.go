package entities

// Product Struct
type Product struct {
	ID       int64 `gorm:"primary_key;AUTO_INCREMENT"`
	Name     string
	Supplier Supplier `gorm:"embedded"`
	Store    Store    `gorm:"embedded"`
}
