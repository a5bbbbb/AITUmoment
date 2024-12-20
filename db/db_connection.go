package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type config struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

type Database struct {
	db *sqlx.DB
}

func newConfig() config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}

	return config{
		host:     getFromEnv("PGHOST", "localhost"),
		user:     getFromEnv("PGUSER", "postgres"),
		port:     getFromEnv("PGPORT", "5432"),
		password: getFromEnv("PGPASSWORD", ""),
		dbname:   getFromEnv("PGDATABASE", "postgres"),
		sslmode:  getFromEnv("PGSSLMODE", "disable"),
	}
}

func getFromEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c config) getDbURLs() (sqlxURL, migrateURL string) {
	sqlxURL = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		c.host,
		c.user,
		c.password,
		c.dbname,
		c.sslmode,
	)

	migrateURL = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.user,
		c.password,
		c.host,
		c.port,
		c.dbname,
		c.sslmode,
	)

	return sqlxURL, migrateURL
}

func getMigrationsPath() (string, error) {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("failed to get caller information")
	}
	return filepath.Join(filepath.Dir(b), "migrations"), nil
}

func NewDatabase() (*Database, error) {
	cfg := newConfig()
	sqlxURL, migrateURL := cfg.getDbURLs()

	db, err := sqlx.Connect("postgres", sqlxURL)

	if err != nil || !isDbConnected(db) {
		return nil, fmt.Errorf("Failed to get database connection, %w", err)
	}

	runMigrations(migrateURL)
	return &Database{db: db}, nil
}

func isDbConnected(db *sqlx.DB) bool {
	err := db.Ping()
	if err != nil {
		return false
	}

	return true
}

func runMigrations(connectionString string) error {
	fmt.Println("START")
	migrationPath, err := getMigrationsPath()
	if err != nil {
		fmt.Println("NO MIGRATION")
		return fmt.Errorf("Failed to get migration path, %w", err)
	}

	sourceURL := "file://" + filepath.ToSlash(migrationPath)
	m, err := migrate.New(sourceURL, connectionString)
	if err != nil {
		fmt.Printf("Migration error: %v\n", err)
		return fmt.Errorf("Could not create migration %w", err)
	}
	defer m.Close()

	version, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			fmt.Println("No migration has been applied yet")
		} else {
			fmt.Printf("Error getting version: %v\n", err)
			return fmt.Errorf("error getting migration version: %w", err)
		}
	} else {
		fmt.Printf("Current version: %d, Dirty: %v\n", version, dirty)
	}

	log.Printf("Current migration version: %d, Dirty: %v", version, dirty)

	// тут сама миграция
	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		fmt.Println(err.Error())
		return fmt.Errorf("Error running migration %w", err)
	}

	newVersion, _, err := m.Version()
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("Error getting new migration version: %w", err)
	}

	if newVersion > version {
		log.Printf("Migrations completed successfully. New version: %d", newVersion)
	} else {
		log.Println("No new migrations to apply")
	}
	fmt.Println("END")

	return nil
}

func (d *Database) GetDB() *sqlx.DB {
	return d.db
}

func (d *Database) Close() error {
	return d.db.Close()
}
