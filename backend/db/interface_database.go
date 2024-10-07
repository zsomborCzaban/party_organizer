package db

type IDatabase interface {
	NewSession()
	AutoMigrate(dst ...interface{}) error

	Create(value interface{}) error
	First(dest interface{}, conds ...interface{}) error
	Find(dest interface{}, conds ...interface{}) error
	Save(value interface{}) error
	Delete(value interface{}, conds ...interface{}) error
	ProcessWhereStatements(conds []QueryParameter)
}
