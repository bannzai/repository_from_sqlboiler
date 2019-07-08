package parser

import "fmt"

type Entity struct {
	FilePath string
	FileReader
}

func (p Entity) Parse() {
	for fieldName, typeName := range parseASTFieldAndType(parseASTFile(p.FilePath), p.extractStructName()) {
		fmt.Printf("structType = %+v\n", fieldName)
		fmt.Printf("typeName = %+v\n", typeName)
	}
}

func (p Entity) extractStructName() string {
	return ""
}
