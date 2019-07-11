package reader

import (
	"path/filepath"
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
	"entitySelectorName":   strutil.Plural,
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
