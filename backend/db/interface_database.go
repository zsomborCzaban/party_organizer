package db

type IDatabase interface {
	NewSession()
	AutoMigrate(dst ...interface{}) error
	//AppendAssociation(entity, associatedEntity interface{}, associationName string) error

	Create(entity interface{}) error
	First(dest interface{}, associations []string, conds ...interface{}) error
	Find(dest interface{}, associations []string, conds ...interface{}) error
	Update(entity interface{}) error
	Delete(entity interface{}) error

	ProcessWhereStatements(conds []QueryParameter)
	Many2ManyQueryId(dest interface{}, associations []string, cond Many2ManyQueryParameter) error
	ReplaceAssociations(param AssociationParameter) error

	TransactionBegin() IDatabase
	TransactionCommit() error
	TransactionRollback() error
}
