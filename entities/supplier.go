package entities

// Supplier Struct
type Supplier struct {
	ID       int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name     string
	Location string
	Phone    string
}
