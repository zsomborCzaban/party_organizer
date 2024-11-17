package db

type IDatabase interface {
	NewSession()
	AutoMigrate(dst ...interface{}) error
	//AppendAssociation(entity, associatedEntity interface{}, associationName string) error

	Create(entity interface{}) error
	First(dest interface{}, conds ...interface{}) error
	Find(dest interface{}, conds ...interface{}) error
	Update(entity interface{}) error
	Save(entity interface{}) error
	Delete(entity interface{}) error
	AddToAssociation(entity interface{}, association string, associatedEntities ...interface{}) error
	DeleteFromAssociation(entity interface{}, association string, associatedEntities ...interface{}) error
	ClearAssociation(entity interface{}, association string) error
	ProcessWhereStatements(conds []QueryParameter)
	Preload(association string)
	Many2ManyQueryId(dest interface{}, cond Many2ManyQueryParameter) error
}
