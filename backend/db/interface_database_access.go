package db

type IDatabaseAccess interface {
	Create(entity interface{}) error
	FindById(id interface{}, associations ...string) (interface{}, error)
	FindAll(associations ...string) (interface{}, error)
	Update(entity interface{}, id uint) error
	Delete(entity interface{}) error

	BatchDelete(conds []QueryParameter) error
	Query(conds []QueryParameter, associations ...string) (interface{}, error)
	//AppendAssociation(entity, associatedEntity interface{}, associationName string) error
	Many2ManyQueryId(cond Many2ManyQueryParameter, associations ...string) (interface{}, error)

	ReplaceAssociations(param AssociationParameter) error
	TransactionBegin() IDatabaseAccess
	TransactionCommit() error
	TransactionRollback() error
}
