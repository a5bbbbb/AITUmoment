package utils

import (
	"fmt"
	"path/filepath"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var migrationPath string

func init() {
	var err error
	migrationPath, err = getMigrationsPath()
	if err != nil {
		logger.GetLogger().Errorf("Failed to get migration path, %w", err)
	}

}

func RunMigrations(connectionString string) error {
	sourceURL := "file://" + filepath.ToSlash(migrationPath)
	m, err := migrate.New(sourceURL, connectionString)
	if err != nil {
		logger.GetLogger().Errorf("Migration error: %v\n", err)
		return fmt.Errorf("Could not create migration %w", err)
	}
	defer m.Close()

	version, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			logger.GetLogger().Infof("No migration has been applied yet")
		} else {
			logger.GetLogger().Errorf("Error getting version: %v\n", err)
			return fmt.Errorf("error getting migration version: %w", err)
		}
	} else {
		logger.GetLogger().Infof("Current version: %d, Dirty: %v\n", version, dirty)
	}

	logger.GetLogger().Infof("Current migration version: %d, Dirty: %v", version, dirty)

	// тут сама миграция
	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		logger.GetLogger().Errorf(err.Error())
		return fmt.Errorf("Error running migration %w", err)
	}

	newVersion, _, err := m.Version()
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return fmt.Errorf("Error getting new migration version: %w", err)
	}

	if newVersion > version {
		logger.GetLogger().Infof("Migrations completed successfully. New version: %d", newVersion)
	} else {
		logger.GetLogger().Info("No new migrations to apply")
	}

	return nil
}

func getMigrationsPath() (string, error) {
	return "./db/migrations", nil
}
