package db

type IDatabaseAccess interface {
	Create(entity interface{}) error
	FindById(id interface{}) (interface{}, error)
	FindAll() (interface{}, error)
	Save(entity interface{}) error
	Update(entity interface{}) error
	Delete(entity interface{}) error
	//DeleteAssociation(entity interface{}, association string) error
	BatchDelete(conds []QueryParameter) error
	Query(conds []QueryParameter) (interface{}, error)
	//AppendAssociation(entity, associatedEntity interface{}, associationName string) error
	Many2ManyQueryId(cond Many2ManyQueryParameter) (interface{}, error)
}
