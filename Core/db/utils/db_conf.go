package utils

import (
	"aitu-moment/logger"
	"aitu-moment/utils"
	"fmt"

)


type DbConfig struct {
    host        string
    port        string
    user        string
    password    string
    dbname      string
    sslmode     string
    sqlxURL     string
    migrateURL  string
}

var Config *DbConfig


func init(){
    Config = &DbConfig{
		host:     utils.GetFromEnv("PGHOST", "localhost"),
		user:     utils.GetFromEnv("PGUSER", "postgres"),
		port:     utils.GetFromEnv("PGPORT", "5432"),
		password: utils.GetFromEnv("PGPASSWORD", ""),
		dbname:   utils.GetFromEnv("PGDATABASE", "postgres"),
		sslmode:  utils.GetFromEnv("PGSSLMODE", "disable"),

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

