package model

type Entity struct {
	Name             string
	PackageName      string
	Fields           []Field
	PrimaryKeys      []PrimaryKey
	ForeignRelations []Entity
	ToManyRelations  []Entity
}
