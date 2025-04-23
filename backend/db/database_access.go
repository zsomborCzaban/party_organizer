package db

import "github.com/rs/zerolog/log"

type DatabaseAccessImpl struct {
	DBEntityProvider IEntityProvider
	DB               IDatabase
}

func NewDatabaseAccessImpl(dbEntityProvider IEntityProvider, db IDatabase) DatabaseAccessImpl {
	return DatabaseAccessImpl{DBEntityProvider: dbEntityProvider, DB: db}
}

func (dbHandler DatabaseAccessImpl) Create(entity interface{}) error {
	dbHandler.DB.NewSession()

	if err := dbHandler.DB.Create(entity); err != nil {
		log.Print(err.Error())
		return NewDBError(err.Error())
	}
	return nil
}

func (dbHandler DatabaseAccessImpl) FindById(id interface{}, associations ...string) (interface{}, error) {
	dbHandler.DB.NewSession()

	result := dbHandler.DBEntityProvider.Create()
	err := dbHandler.DB.First(result, associations, id)
	if err != nil {
		log.Print(err.Error())
		return nil, NewDBError(err.Error())
	}
	return result, nil
}

func (dbHandler DatabaseAccessImpl) FindAll(associations ...string) (interface{}, error) {
	dbHandler.DB.NewSession()

	entities := dbHandler.DBEntityProvider.CreateArray()
	err := dbHandler.DB.Find(entities, associations)
	if err != nil {
		log.Print(err.Error())
		return entities, NewDBError(err.Error())
	}
	return entities, nil
}

func (dbHandler DatabaseAccessImpl) Update(entity interface{}, id uint) error {
	dbHandler.DB.NewSession()

	if err := dbHandler.DB.Update(entity, dbHandler.DBEntityProvider.Create(), id); err != nil {
		log.Print(err.Error())
		return NewDBError(err.Error())
	}
	return nil
}

func (dbHandler DatabaseAccessImpl) Delete(entity interface{}) error {
	dbHandler.DB.NewSession()

	err := dbHandler.DB.Delete(entity)
	if err != nil {
		log.Print(err.Error())
		return NewDBError(err.Error())
	}
	return nil
}

func (dbHandler DatabaseAccessImpl) BatchDelete(conds []QueryParameter) error {
	dbHandler.DB.NewSession()
	dbHandler.DB.ProcessWhereStatements(conds)

	entity := dbHandler.DBEntityProvider.Create()
	err := dbHandler.DB.Delete(entity)
	if err != nil {
		log.Print(err.Error())
		return NewDBError(err.Error())
	}
	return nil
}

func (dbHandler DatabaseAccessImpl) Query(conds []QueryParameter, associations ...string) (interface{}, error) {
	dbHandler.DB.NewSession()
	dbHandler.DB.ProcessWhereStatements(conds)

	entities := dbHandler.DBEntityProvider.CreateArray()
	err := dbHandler.DB.Find(entities, associations)
	if err != nil {
		log.Print(err.Error())

		return entities, NewDBError(err.Error())
	}
	return entities, nil
}

//func (dbHandler DatabaseAccessImpl) AppendAssociation(entity, associatedEntity interface{}, associationName string) error {
//	dbHandler.DB.NewSession()
//	return dbHandler.DB.AppendAssociation(&entity, &associatedEntity, associationName)
//}

func (dbHandler DatabaseAccessImpl) Many2ManyQueryId(cond Many2ManyQueryParameter, associations ...string) (interface{}, error) {
	dbHandler.DB.NewSession()

	entities := dbHandler.DBEntityProvider.CreateArray()
	err := dbHandler.DB.Many2ManyQueryId(entities, associations, cond)
	if err != nil {
		log.Print(err.Error())
		return entities, NewDBError(err.Error())
	}
	return entities, nil
}

func (dbHandler DatabaseAccessImpl) ReplaceAssociations(param AssociationParameter) error {
	dbHandler.DB.NewSession()
	return dbHandler.DB.ReplaceAssociations(param)
}

func (dbHandler DatabaseAccessImpl) TransactionBegin() IDatabaseAccess {
	return NewDatabaseAccessImpl(dbHandler.DBEntityProvider, dbHandler.DB.TransactionBegin())
}

func (dbHandler DatabaseAccessImpl) TransactionCommit() error {
	return dbHandler.DB.TransactionCommit()
}

func (dbHandler DatabaseAccessImpl) TransactionRollback() error {
	return dbHandler.DB.TransactionRollback()

}
