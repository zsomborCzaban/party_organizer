package db

type IDatabaseAccessManager interface {
	RegisterEntity(name string, dbEntityProvider IEntityProvider) IDatabaseAccess
	GetRegisteredDBAccess(name string) IDatabaseAccess
	Close() error
}
