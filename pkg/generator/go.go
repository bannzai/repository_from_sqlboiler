package generator

import "github.com/bannzai/repository_from_sqlboiler/pkg/writer"

type GoCodeGenerator struct {
	SourceCodeReader
	GoFormatter
	Writer writer.FileWriter
}

func (generator GoCodeGenerator) Generate() {
	sourceCode := generator.SourceCodeReader.Read()

}
