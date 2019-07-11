package generator

import "text/template"

type TemplateReader interface {
	Read() *template.Template
}
