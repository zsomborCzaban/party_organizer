package db

type QueryParameter struct {
	Field    string
	Operator string
	Value    interface{}
}

type AssociationParameter struct {
	Model       interface{}
	Association string
	Values      interface{}
}
