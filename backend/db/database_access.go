package db

type DatabaseAccessImpl struct {
	DBEntityProvider IEntityProvider
	DB               IDatabase
}

func NewDatabaseAccessImpl(dbEntityProvider IEntityProvider, db IDatabase) DatabaseAccessImpl {
	return DatabaseAccessImpl{DBEntityProvider: dbEntityProvider, DB: db}
}

func (dbHandler DatabaseAccessImpl) Create(entity interface{}) error {
	dbHandler.DB.NewSession()

	err := dbHandler.DB.Create(entity)
	if err != nil {
		return NewDBError(err.Error())
	}
	return nil
}

func (dbHandler DatabaseAccessImpl) FindById(id interface{}) (interface{}, error) {
	dbHandler.DB.NewSession()

	result := dbHandler.DBEntityProvider.Create()
	err := dbHandler.DB.First(result, id)
	if err != nil {
		return nil, NewDBError(err.Error())
	}
	return result, nil
}

func (dbHandler DatabaseAccessImpl) Update(entity interface{}) error {
	dbHandler.DB.NewSession()

	err := dbHandler.DB.Save(entity)
	if err != nil {
		return NewDBError(err.Error())
	}
	return nil
}

func (dbHandler DatabaseAccessImpl) FindAll() (interface{}, error) {
	dbHandler.DB.NewSession()

	entities := dbHandler.DBEntityProvider.CreateArray()
	err := dbHandler.DB.Find(entities)
	if err != nil {
		return entities, NewDBError(err.Error())
	}
	return entities, nil
}

func (dbHandler DatabaseAccessImpl) Delete(entity interface{}) error {
	dbHandler.DB.NewSession()

	err := dbHandler.DB.Delete(entity)
	if err != nil {
		return NewDBError(err.Error())
	}
	return nil
}

func (dbHandler DatabaseAccessImpl) Query(conds []QueryParameter) (interface{}, error) {
	dbHandler.DB.NewSession()
	dbHandler.DB.ProcessWhereStatements(conds)

	entities := dbHandler.DBEntityProvider.CreateArray()
	err := dbHandler.DB.Find(entities)
	if err != nil {
		return entities, NewDBError(err.Error())
	}
	return entities, nil
}

//func (dbHandler DatabaseAccessImpl) AppendAssociation(entity, associatedEntity interface{}, associationName string) error {
//	dbHandler.DB.NewSession()
//	return dbHandler.DB.AppendAssociation(&entity, &associatedEntity, associationName)
//}
