package db

type IEntityProvider interface {
	Create() interface{}
	CreateArray() interface{}
}
