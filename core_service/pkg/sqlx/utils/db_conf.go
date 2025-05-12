package utils

import (
	"fmt"
	"os"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
)

type DbConfig struct {
	host       string
	port       string
	user       string
	password   string
	dbname     string
	sslmode    string
	sqlxURL    string
	migrateURL string
}

var Config *DbConfig

func GetFromEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func init() {
	Config = &DbConfig{
		host:     GetFromEnv("PGHOST", "localhost"),
		user:     GetFromEnv("PGUSER", "postgres"),
		port:     GetFromEnv("PGPORT", "5432"),
		password: GetFromEnv("PGPASSWORD", ""),
		dbname:   GetFromEnv("PGDATABASE", "postgres"),
		sslmode:  GetFromEnv("PGSSLMODE", "disable"),
	}

	logger.GetLogger().Info("Successfully loaded DB config")

}

func (c DbConfig) GetDbURLs() (sqlxURL, migrateURL string) {

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
