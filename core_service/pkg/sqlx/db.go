package sqlx

import (
	"sync"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/core_service/pkg/sqlx/utils"

	"github.com/jmoiron/sqlx"
)

var (
	instance *sqlx.DB
	once     sync.Once
)

func init() {
	logger.GetLogger().Info("Connecting to DB...")
	once.Do(startUpDB)
}

func startUpDB() {
	cfg := utils.Config
	sqlxURL, migrateURL := cfg.GetDbURLs()

	db, err := sqlx.Connect("postgres", sqlxURL)

	//Here you can also set up the db configs like:
	// db.SetMaxOpenConns(23)

	if err != nil || !isDbConnected(db) {
		logger.GetLogger().Errorf("Could not establish the connection to the DB! %v", err.Error())
	}

	instance = db

	err = utils.RunMigrations(migrateURL)

	if err != nil {
		logger.GetLogger().Errorf("Could not run DB migrations! %v", err.Error())
	}

	logger.GetLogger().Info("Successfully established connection with DB")

}

func isDbConnected(db *sqlx.DB) bool {
	err := db.Ping()
	if err != nil {
		return false
	}

	return true
}

func GetDB() *sqlx.DB {
	logger.GetLogger().Trace("Getting DB connection")
	return instance
}

func Close() error {
	logger.GetLogger().Info("Closing DB connection....")
	err := instance.Close()

	if err != nil {
		logger.GetLogger().Errorf("Error during closing DB connection, %v", err.Error())
	}

	return err
}
