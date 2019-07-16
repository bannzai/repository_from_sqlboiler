package reader

import (
	"path/filepath"
	"strings"
	"text/template"

	"github.com/bannzai/repository_from_sqlboiler/pkg/strutil"
)

type Template struct {
	FilePath string
}

var functions = template.FuncMap{
	"sqlParameterCase":     strutil.SnakeCase,
	"golangArgumentCase":   golangArgumentCase,
	"golangVariableCase":   golangVariableCase,
	"golangStructNameCase": strutil.SpecializeUpperCamelCase,
	"plural":               strutil.Plural,
	"entitySelectorName":   entitySelectorName,
}

func entitySelectorName(str string) string {
	str = strings.ToLower(str)
	str = strutil.SnakeCase(str)
	str = strutil.Plural(str)
	str = strutil.UpperCamelCase(str)
	return str
}

func golangArgumentCase(str string) string {
	return golangVariableCase(str)
}

func golangVariableCase(str string) string {
	str = strutil.SpecializeLowerCamelCase(str)
	str = strutil.EscapedReservedWord(str)
	return str
}

func (t Template) Read() *template.Template {
	filePath := t.FilePath
	base := filepath.Base(filePath)
	return template.Must(template.New(base).Funcs(functions).ParseFiles(filePath))
}
