package parser

import (
	"github.com/bannzai/repository_from_sqlboiler/pkg/model"
)

type Entity struct {
	EntityName string
	FilePath   string
	FileReader
}

func (p Entity) Parse() model.Entity {
	return model.Entity{
		Name:        p.EntityName,
		Fields:      p.mapFields(),
		PrimaryKeys: p.mapPrimaryKeys(),
	}
}

func (p Entity) mapPrimaryKeys() []model.PrimaryKey {
	primaryKeys := []model.PrimaryKey{}
	for _, primaryKeyName := range parseASTFieldAndType(parseASTFile(p.FilePath), p.extractStructName()) {
		primaryKeys = append(
			primaryKeys,
			model.PrimaryKey{
				Name: primaryKeyName,
			})
	}
	return primaryKeys
}

func (p Entity) mapFields() []model.Field {
	fields := []model.Field{}
	for fieldName, typeName := range parseASTFieldAndType(parseASTFile(p.FilePath), p.extractStructName()) {
		fields = append(fields, model.Field{
			Name:     fieldName,
			TypeName: typeName,
		})
	}
	return fields
}
