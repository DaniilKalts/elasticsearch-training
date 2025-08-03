package database

import (
	"database/sql"
	"fmt"
	"github.com/DaniilKalts/elasticsearch-training/internal/application"
	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func InitDB(cfg *application.AppConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	migrationsDir := "/db/migrations"
	if err := initMigrations(db, migrationsDir); err != nil {
		return nil, err
	}

	return db, err
}

func initMigrations(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}
	if err := goose.Up(db, dir); err != nil {
		return err
	}

	return nil
}
