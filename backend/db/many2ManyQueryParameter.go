package db

type Many2ManyQueryParameter struct {
	QueriedTable            string
	Many2ManyTable          string
	M2MQueriedColumnName    string
	M2MConditionColumnName  string
	M2MConditionColumnValue interface{}
	OrActive                bool
	OrConditionColumnName   string
	OrConditionColumnValue  interface{}
}
