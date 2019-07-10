package reader

import (
	"text/template"

	"github.com/bannzai/repository_from_sqlboiler/pkg/strutil"
)

type Template struct {
	FilePath string
}

var functions = template.FuncMap{
	"sqlParameterCase":     strutil.SnakeCase,
	"golangArgumentCase":   strutil.SpecializeLowerCamelCase,
	"golangVariableCase":   strutil.SpecializeLowerCamelCase,
	"golangStructNameCase": strutil.SpecializeUpperCamelCase,
}

func (t Template) Read() *template.Template {
	filePath := t.FilePath
	return template.Must(template.New(filePath).Funcs(functions).ParseFiles(filePath))
}
