package parser

import (
	"github.com/bannzai/repository_from_sqlboiler/pkg/model"
)

type Entity struct {
	FilePath string
}

func (p Entity) Parse() model.Entity {
	return model.Entity{
		Name:        p.extractEntityName(),
		Fields:      p.mapFields(),
		PrimaryKeys: p.mapPrimaryKeys(),
	}
}

func (p Entity) mapPrimaryKeys() []model.PrimaryKey {
	primaryKeys := []model.PrimaryKey{}
	for _, primaryKeyName := range parseASTPrimaryKeyFields(parseASTFile(p.FilePath)) {
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
	for fieldName, typeName := range parseASTFieldAndType(parseASTFile(p.FilePath), p.extractEntityName()) {
		fields = append(fields, model.Field{
			Name:     fieldName,
			TypeName: typeName,
		})
	}
	return fields
}

func (e Entity) extractEntityName() string {
	return parseASTBaseStructName(parseASTFile(e.FilePath))
}
