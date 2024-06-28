package database

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/koanf/v2"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB(k *koanf.Koanf, logger *slog.Logger) error {

	// Construct connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		k.String("db.user"), k.String("db.password"), k.String("db.dbname"))

	// Attempt to connect to PostgreSQL database
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	// Assign the connected DB to the package-level variable
	DB = db

	log.Println("Connected to PostgreSQL database successfully")
	return nil
}
