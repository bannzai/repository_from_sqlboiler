package generator

import (
	"fmt"
	"strings"

	"github.com/bannzai/repository_from_sqlboiler/pkg/model"
	"github.com/bannzai/repository_from_sqlboiler/pkg/strutil"
)

func fetchByPrimaryKey(entity model.Entity) string {
	return fmt.Sprintf(
		"%v(%v) entity.%v",
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
			// Remove _ and same character size.
			lhs := strings.ToLower(strutil.LowerCamelCase(column.Name))
			rhs := strings.ToLower(strutil.LowerCamelCase(primaryKey.Name))
			if lhs == rhs {
				typeName = column.TypeName
			}
		}
		return typeName
	}

	condition := "ctx Context, "
	for i, primaryKey := range entity.PrimaryKeys {
		if i > 0 {
			condition += ", "
		}
		column := strutil.SpecializeLowerCamelCase(primaryKey.Name)
		column = strutil.EscapedReservedWord(column)
		typeName := findPrimaryKeyType(primaryKey, entity)
		condition += fmt.Sprintf("%v %v", column, typeName)
	}

	return condition
}

func listOfPrimaryKeys(entity model.Entity) string {
	content := ""
	for i, primaryKey := range entity.PrimaryKeys {
		if i > 0 {
			content += ", "
		}
		content += strutil.EscapedReservedWord(primaryKey.Name)
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
		variable := strutil.LowerCamelCase(primaryKey.Name)
		variable = strutil.EscapedReservedWord(variable)
		condition += fmt.Sprintf("qm.Where(\"%v=?\", %v)", column, variable)
	}

	return condition
}
