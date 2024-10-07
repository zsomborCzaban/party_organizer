package db

type IDatabaseAccess interface {
	Create(entity interface{}) error
	FindById(id interface{}) (interface{}, error)
	FindAll() (interface{}, error)
	Update(entity interface{}) error
	Delete(entity interface{}) error

	Query(conds []QueryParameter) (interface{}, error)
	//AppendAssociation(entity, associatedEntity interface{}, associationName string) error
}
