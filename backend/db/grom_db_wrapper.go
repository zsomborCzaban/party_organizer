package db

import (
	"fmt"
	"gorm.io/gorm"
)

type GormDBWrapper struct {
	DB *gorm.DB
}

var COLUMNS_TO_OMIT_DURING_UPDATE = []string{"created_at", "deleted_at"}

func NewGormDBWrapper(db *gorm.DB) *GormDBWrapper {
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

func (dbWrapper *GormDBWrapper) First(dest interface{}, associations []string, conds ...interface{}) error {
	db := dbWrapper.DB.Session(&gorm.Session{NewDB: true})
	for _, association := range associations {
		db = db.Preload(association)
	}
	return db.First(dest, conds).Error
}

func (dbWrapper *GormDBWrapper) Find(dest interface{}, associations []string, conds ...interface{}) error {
	//db := dbWrapper.DB.Session(&gorm.Session{NewDB: true})
	for _, association := range associations {
		dbWrapper.DB = dbWrapper.DB.Preload(association)
	}

	//dbWrapper.DB.Error = nil
	return dbWrapper.DB.Find(dest, conds).Error //causes concurrent map writes once
}

func (dbWrapper *GormDBWrapper) Update(entity interface{}, model interface{}, id uint) error {
	return dbWrapper.DB.Model(model).Where("id = ?", id).Omit(COLUMNS_TO_OMIT_DURING_UPDATE...).Updates(entity).Error //caused concurent mapp writes once
}

func (dbWrapper *GormDBWrapper) Delete(entity interface{}) error {
	return dbWrapper.DB.Delete(entity).Error
}

//func (dbWrapper *GormDBWrapper) AddToAssociation(entity interface{}, association string, associatedEntities ...interface{}) error {
//	return dbWrapper.DB.Model(entity).Association(association).Append(associatedEntities[0])
//}
//
//func (dbWrapper *GormDBWrapper) DeleteFromAssociation(entity interface{}, association string, associatedEntities ...interface{}) error {
//	return dbWrapper.DB.Model(entity).Association(association).Delete(associatedEntities)
//}

//func (dbWrapper *GormDBWrapper) ClearAssociation(entity interface{}, association string) error {
//	return dbWrapper.DB.Model(entity).Association(association).Clear()
//}

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

//func (dbWrapper *GormDBWrapper) AppendAssociation(entity, associatedEntity interface{}, associationName string) error {
//	association := dbWrapper.DB.Model(entity).Association(associationName)
//	err := association.Append(associatedEntity)
//	return err
//}

func (dbWrapper *GormDBWrapper) Many2ManyQueryId(dest interface{}, associations []string, cond Many2ManyQueryParameter) error {
	db := dbWrapper.DB.Session(&gorm.Session{NewDB: true})
	if !cond.OrActive {
		query := fmt.Sprintf(
			"SELECT * FROM %s WHERE id IN (SELECT %s FROM %s WHERE %s = ?)",
			cond.QueriedTable, cond.M2MQueriedColumnName, cond.Many2ManyTable, cond.M2MConditionColumnName,
		)
		for _, preloadColumn := range associations {
			db = db.Preload(preloadColumn)
		}
		return db.Raw(query, cond.M2MConditionColumnValue).Find(dest).Error
	} else {
		query := fmt.Sprintf(
			"SELECT * FROM %s WHERE id IN (SELECT %s FROM %s WHERE %s = ? OR %s = ?)",
			cond.QueriedTable, cond.M2MQueriedColumnName, cond.Many2ManyTable, cond.M2MConditionColumnName, cond.OrConditionColumnName,
		)
		for _, preloadColumn := range associations {
			db = db.Preload(preloadColumn)
		}
		return db.Raw(query, cond.M2MConditionColumnValue, cond.OrConditionColumnValue).Find(dest).Error
	}
}

func (dbWrapper *GormDBWrapper) ReplaceAssociations(param AssociationParameter) error {
	return dbWrapper.DB.Model(param.Model).Association(param.Association).Replace(param.Values)
}

func (dbWrapper *GormDBWrapper) TransactionBegin() IDatabase {
	tr := GormDBWrapper{
		DB: dbWrapper.DB.Begin(),
	}
	return &tr
}

func (dbWrapper *GormDBWrapper) TransactionCommit() error {
	return dbWrapper.DB.Commit().Error
}

func (dbWrapper *GormDBWrapper) TransactionRollback() error {
	return dbWrapper.DB.Rollback().Error

}
