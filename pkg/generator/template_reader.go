package generator

import "text/template"

type TemplateReader interface {
	Read(filePath string) *template.Template
}
