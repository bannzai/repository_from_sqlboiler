package generator

import (
	"bytes"
	"fmt"
	"strings"

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
		"Entity":            entity,
		"FetchByPrimaryKey": fetchByPrimaryKey(entity),
		"SQLQueryArguments": sqlQueryArguments(entity),
		"PrimaryKeys":       listOfPrimaryKeys(entity),
	}); err != nil {
		panic(err)
	}
	return buf.String()
}

func fetchByPrimaryKey(entity model.Entity) string {
	return fmt.Sprintf(
		"%v(%v) %v",
		fetchByPrimaryKeyFunctionName(entity),
		fetchByPrimaryKeyFunctionArgument(entity),
		entity.Name,
	)
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
			if strings.ToLower(column.Name) == strings.ToLower(primaryKey.Name) {
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

func listOfPrimaryKeys(entity model.Entity) string {
	content := ""
	for _, primaryKey := range entity.PrimaryKeys {
		content += primaryKey.Name
	}
	return content
}

func sqlQueryArguments(entity model.Entity) string {
	condition := ""
	for i, primaryKey := range entity.PrimaryKeys {
		if i > 0 {
			condition += ", "
		}
		column := strutil.SnakeCase(primaryKey.Name)
		condition += fmt.Sprintf("qm.Where(\"%v=?\", %v)", column, column)
	}

	return condition
}
