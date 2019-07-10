package reader

import (
	"text/template"

	"github.com/bannzai/repository_from_sqlboiler/pkg/strutil"
)

type Template struct{}

var functions = template.FuncMap{
	"sqlParameterCase":     strutil.SnakeCase,
	"golangArgumentCase":   strutil.SpecializeToLowerCamelCase,
	"golangStructNameCase": strutil.SpecializeUpperCamelCase,
}

func (Template) Read(filePath string) *template.Template {
	return template.Must(template.New(filePath).Funcs(functions).ParseFiles(filePath))
}
