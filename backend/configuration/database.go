package configuration

import (
	"github.com/zsomborCzaban/party_organizer/common"
	"github.com/zsomborCzaban/party_organizer/db"
	"gorm.io/gorm/logger"
	log2 "log"
	"os"
	"time"
)

type DbAccessManagerBuilderFunc func(dbConnectionUrl string, dbLogger logger.Interface) db.IDatabaseAccessManager

func CreateDbAccessManager(builderfunc DbAccessManagerBuilderFunc) db.IDatabaseAccessManager {
	dbLogger := logger.New(
		log2.New(os.Stdout, "\r\n", log2.LstdFlags), //io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	return builderfunc(common.DEFAULT_LOCAL_DB_NAME, dbLogger)
}
