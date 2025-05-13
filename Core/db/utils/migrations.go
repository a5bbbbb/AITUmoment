package utils

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"path/filepath"
)

var migrationPath string

func init() {
	var err error
	migrationPath, err = getMigrationsPath()
	if err != nil {
		log.Print("Failed to get migration path, %w", err)
	}

}

func RunMigrations(connectionString string) error {
	sourceURL := "file://" + filepath.ToSlash(migrationPath)
	fmt.Println(sourceURL)
	fmt.Println("YOU")
	fmt.Println(connectionString)
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

func getMigrationsPath() (string, error) {
	return "./db/migrations", nil
}
