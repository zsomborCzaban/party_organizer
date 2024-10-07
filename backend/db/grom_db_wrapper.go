package db

import (
	"fmt"
	"gorm.io/gorm"
)

type GormDBWrapper struct {
	DB *gorm.DB
}

func NewGormDBWrapper(dbEntity interface{}, db *gorm.DB) *GormDBWrapper {
	return &GormDBWrapper{DB: db}
}

func (dbWrapper *GormDBWrapper) NewSession() {
	dbWrapper.DB = dbWrapper.DB.Session(&gorm.Session{NewDB: true})
}

func (dbWrapper *GormDBWrapper) AutoMigrate(dst ...interface{}) error {
	return dbWrapper.DB.AutoMigrate(dst)
}

func (dbWrapper *GormDBWrapper) Create(entity interface{}) error {
	return dbWrapper.DB.Create(entity).Error
}

func (dbWrapper *GormDBWrapper) First(dest interface{}, conds ...interface{}) error {
	return dbWrapper.DB.First(dest, conds).Error
}

func (dbWrapper *GormDBWrapper) Save(entity interface{}) error {
	return dbWrapper.DB.Save(entity).Error
}

func (dbWrapper *GormDBWrapper) Find(dest interface{}, conds ...interface{}) error {
	return dbWrapper.DB.Find(dest, conds).Error
}

func (dbWrapper *GormDBWrapper) Delete(value interface{}, conds ...interface{}) error {
	return dbWrapper.DB.Delete(value, conds).Error
}

func (dbWrapper *GormDBWrapper) ProcessWhereStatements(conds []QueryParameter) {
	for _, queryParam := range conds {
		if len(queryParam.Operator) > 0 {
			format := fmt.Sprintf("%s %s ?", queryParam.Field, queryParam.Operator)
			dbWrapper.DB = dbWrapper.DB.Where(format, queryParam.Value)
		} else {
			dbWrapper.DB = dbWrapper.DB.Where(queryParam.Value)
		}
	}
}
