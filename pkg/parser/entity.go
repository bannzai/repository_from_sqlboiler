package parser

import "fmt"

type Entity struct {
	FilePath string
	FileReader
}

func (p Entity) Parse() {
	for typeName, structType := range parseASTStructs(parseASTFile(p.FilePath)) {
		fmt.Printf("typeName = %+v\n", typeName)
		fmt.Printf("structType = %+v\n", structType)
	}
}
