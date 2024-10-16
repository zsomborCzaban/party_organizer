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

func (dbWrapper *GormDBWrapper) Update(entity interface{}) error {
	return dbWrapper.DB.Model(entity).Omit("id", "created_at", "deleted_at").Updates(entity).Error
}

func (dbWrapper *GormDBWrapper) Save(entity interface{}) error {
	return dbWrapper.DB.Omit("created_at", "deleted_at").Save(entity).Error
}

func (dbWrapper *GormDBWrapper) Find(dest interface{}, conds ...interface{}) error {
	return dbWrapper.DB.Find(dest, conds).Error
}

func (dbWrapper *GormDBWrapper) Delete(entity interface{}, conds ...interface{}) error {
	return dbWrapper.DB.Delete(entity, conds).Error
}

//func (dbWrapper *GormDBWrapper) DeleteAssociation(entity interface{}, association string) error {
//	return dbWrapper.DB.Model(entity).Association(association).Unscoped().Clear()
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

func (dbWrapper *GormDBWrapper) Many2ManyQueryId(dest interface{}, cond Many2ManyQueryParameter) error {
	if !cond.OrActive {
		query := fmt.Sprintf(
			"SELECT * FROM %s WHERE id IN (SELECT %s FROM %s WHERE %s = ?)",
			cond.QueriedTable, cond.M2MQueriedColumnName, cond.Many2ManyTable, cond.M2MConditionColumnName,
		)
		return dbWrapper.DB.Raw(query, cond.M2MConditionColumnValue).Scan(dest).Error
	} else {
		query := fmt.Sprintf(
			"SELECT * FROM %s WHERE id IN (SELECT %s FROM %s WHERE %s = ? OR %s = ?)",
			cond.QueriedTable, cond.M2MQueriedColumnName, cond.Many2ManyTable, cond.M2MConditionColumnName, cond.OrConditionColumnName,
		)
		return dbWrapper.DB.Raw(query, cond.M2MConditionColumnValue, cond.OrConditionColumnValue).Scan(dest).Error
	}
	//return dbWrapper.DB.Model(model).Preload(preload, "id = ?", 3).Find(dest).Error
}
