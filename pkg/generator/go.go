package generator

import (
	"bytes"
	"fmt"

	"github.com/bannzai/repository_from_sqlboiler/pkg/model"
	"github.com/bannzai/repository_from_sqlboiler/pkg/strutil"
	"github.com/bannzai/repository_from_sqlboiler/pkg/writer"
)

type GoCodeGenerator struct {
	DestinationFilePath string
	EntityParser
	GoFormatter
	TemplateReader
	Writer writer.FileWriter
}

func (generator GoCodeGenerator) Generate() {
	entity := generator.EntityParser.Parse()
	content := generator.content(entity)
	generator.Writer.Write(content)
	generator.GoImports()
	generator.GoFormat()
}

func (generator GoCodeGenerator) content(entity model.Entity) string {
	buf := &bytes.Buffer{}
	if err := generator.TemplateReader.Read().Execute(buf, map[string]interface{}{
		"Entity":                            entity,
		"FetchByPrimaryKeyFunctionName":     fetchByPrimaryKeyFunctionName(entity),
		"FetchByPrimaryKeyFunctionArgument": fetchByPrimaryKeyFunctionArgument(entity),
		"SQLQueryArguments":                 sqlQueryArguments(entity),
	}); err != nil {
		panic(err)
	}
	return buf.String()
}

func fetchByPrimaryKeyFunctionName(entity model.Entity) string {
	const prefix = "FetchBy"
	suffix := ""
	for i, primaryKey := range entity.PrimaryKeys {
		if i > 0 {
			suffix += "And"
		}
		suffix += strutil.SpecializeUpperCamelCase(primaryKey.Name)
	}
	return prefix + suffix
}

func fetchByPrimaryKeyFunctionArgument(entity model.Entity) string {
	findPrimaryKeyType := func(primaryKey model.PrimaryKey, entity model.Entity) string {
		typeName := ""
		for _, column := range entity.Fields {
			if column.Name == primaryKey.Name {
				typeName = column.TypeName
			}
		}
		return typeName
	}

	condition := ""
	for i, primaryKey := range entity.PrimaryKeys {
		if i > 0 {
			condition += ", "
		}
		column := strutil.SpecializeLowerCamelCase(primaryKey.Name)
		typeName := findPrimaryKeyType(primaryKey, entity)
		condition += fmt.Sprintf("%v %v", column, typeName)
	}

	return condition
}

func sqlQueryArguments(entity model.Entity) string {

	condition := ""
	for i, primaryKey := range entity.PrimaryKeys {
		if i > 0 {
			condition += ", "
		}
		column := strutil.SnakeCase(primaryKey.Name)
		condition += fmt.Sprintf("qm.Where(\"%v\"=?, %v)", column, column)
	}

	return condition
}
