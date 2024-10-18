package db

type IDatabaseAccess interface {
	Create(entity interface{}) error
	FindById(id interface{}, associations ...string) (interface{}, error)
	FindAll() (interface{}, error)
	Save(entity interface{}) error
	Update(entity interface{}) error
	Delete(entity interface{}) error
	DeleteFromAssociation(entity interface{}, association string, associatedEntities ...interface{}) error
	AddToAssociation(entity interface{}, association string, associatedEntities ...interface{}) error
	BatchDelete(conds []QueryParameter) error
	Query(conds []QueryParameter) (interface{}, error)
	//AppendAssociation(entity, associatedEntity interface{}, associationName string) error
	Many2ManyQueryId(cond Many2ManyQueryParameter) (interface{}, error)
}
