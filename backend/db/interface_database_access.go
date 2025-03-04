package db

type IDatabaseAccess interface {
	Create(entity interface{}) error
	FindById(id interface{}, associations ...string) (interface{}, error)
	FindAll(associations ...string) (interface{}, error)
	Update(entity interface{}) error
	Delete(entity interface{}) error

	AddToAssociation(entity interface{}, association string, associatedEntities ...interface{}) error
	DeleteFromAssociation(entity interface{}, association string, associatedEntities ...interface{}) error
	ClearAssociation(entity interface{}, associations ...string) error
	BatchDelete(conds []QueryParameter) error
	Query(conds []QueryParameter, associations ...string) (interface{}, error)
	//AppendAssociation(entity, associatedEntity interface{}, associationName string) error
	Many2ManyQueryId(cond Many2ManyQueryParameter, associations ...string) (interface{}, error)
}
