package database

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/knadh/koanf/v2"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(k *koanf.Koanf, logger *slog.Logger) error {

	// Construct connection string
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		k.String("db.host"), k.String("db.user"), k.String("db.password"), k.String("db.dbname"), k.String("db.port"))

	// Attempt to connect to PostgreSQL database
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	// Assign the connected DB to the package-level variable
	DB = db

	log.Println("Connected to PostgreSQL database successfully")
	// Migrate models
	DB.AutoMigrate(&User{},
		&Member{},
		&CoreMember{},
		&Staff{},
		&Event{},
		&Attendee{},
		&Comment{},
		&Project{},
		&ProjectRequest{},
		&ProjectMember{},
		&Notice{})
	return nil
}
