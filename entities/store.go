package entities

// Store Struct
type Store struct {
	ID       int64 `gorm:"primary_key;AUTO_INCREMENT"`
	Name     string
	Location string
}
