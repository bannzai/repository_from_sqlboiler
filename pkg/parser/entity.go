package parser

import (
	"path/filepath"

	"github.com/bannzai/repository_from_sqlboiler/pkg/model"
	"github.com/bannzai/repository_from_sqlboiler/pkg/strutil"
)

type Entity struct {
	FilePath string
	FileReader
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
	for _, primaryKeyName := range parseASTFieldAndType(parseASTFile(p.FilePath), p.extractEntityName()) {
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
	filename := filepath.Base(e.FilePath)
	extension := filepath.Ext(e.FilePath)
	snakeCaseEntityName := filename[0 : len(filename)-len(extension)]
	entityName := strutil.UpperCamelCase(snakeCaseEntityName)
	return entityName
}
