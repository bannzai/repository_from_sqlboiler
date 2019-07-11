package model

type Entity struct {
	Name        string
	Fields      []Field
	PrimaryKeys []PrimaryKey
}
