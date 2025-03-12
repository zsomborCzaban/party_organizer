package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDatabaseAccessManager struct {
	DB               *gorm.DB
	DBAccessRegistry map[string]IDatabaseAccess
}

func CreateGormDatabaseAccessManager(dbConnectionUrl string, newLogger logger.Interface) IDatabaseAccessManager {
	db, err := gorm.Open(sqlite.Open(dbConnectionUrl+"?_foreign_keys=on"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	return GormDatabaseAccessManager{DB: db, DBAccessRegistry: make(map[string]IDatabaseAccess)}
}

func (dam GormDatabaseAccessManager) RegisterEntity(name string, dbEntityProvider IEntityProvider) IDatabaseAccess {
	_, ok := dam.DBAccessRegistry[name]
	if ok {
		panic("DB access is already registered for name: " + name)
	}

	entity := dbEntityProvider.Create()
	migrateErr := dam.DB.AutoMigrate(entity)

	if migrateErr != nil {
		panic("failed to migrate database: " + name + " ; error: " + migrateErr.Error())
	}

	dbWrapper := NewGormDBWrapper(dam.DB)
	access := NewDatabaseAccessImpl(dbEntityProvider, dbWrapper)

	dam.DBAccessRegistry[name] = access

	return access
}

func (dam GormDatabaseAccessManager) GetRegisteredDBAccess(name string) IDatabaseAccess {
	return dam.DBAccessRegistry[name]
}

func (dam GormDatabaseAccessManager) Close() error {
	sqlDB, err := dam.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
