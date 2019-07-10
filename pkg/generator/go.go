package generator

import (
	"bytes"

	"github.com/bannzai/repository_from_sqlboiler/pkg/model"
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
	if err := generator.TemplateReader.Read(generator.DestinationFilePath).Execute(buf, map[string]interface{}{
		"Entity": entity,
	}); err != nil {
		panic(err)
	}
	return buf.String()
}
