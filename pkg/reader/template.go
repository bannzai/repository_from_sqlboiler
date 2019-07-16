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

// reference: https://github.com/volatiletech/sqlboiler/blob/4dea9a5f77e3ec6090ad455df3d408e47e077700/strmangle/strmangle.go#L25
var lowercaseWords = map[string]string{
	"ACL":   "Acl",
	"API":   "Api",
	"ASCII": "Ascii",
	"CPU":   "Cpu",
	"EOF":   "Eof",
	"GUID":  "Guid",
	"ID":    "Id",
	"IP":    "Ip",
	"JSON":  "Json",
	"RAM":   "Ram",
	"SLA":   "Sla",
	"UDP":   "Udp",
	"UI":    "Ui",
	"UID":   "Uid",
	"UUID":  "Uuid",
	"URI":   "Uri",
	"URL":   "Url",
	"UTF8":  "Utf8",
}

func entitySelectorName(str string) string {
	for keyword, converted := range lowercaseWords {
		if strings.Contains(str, keyword) {
			str = strings.ReplaceAll(str, keyword, converted)
		}
	}
	str = strutil.Plural(str)
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
