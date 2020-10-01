package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gotraining/entities"
)

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&entities.Supplier{})
	db.AutoMigrate(&entities.Store{})
	db.AutoMigrate(&entities.Product{})
}

// OpenDB func
func OpenDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")

	fmt.Println(host)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", host, port, user, name, ssl, password)
	conf := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	})
	gormDB, err := gorm.Open(conf, &gorm.Config{})

	migrateDB(gormDB)

	return gormDB
}
