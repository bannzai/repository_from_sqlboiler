package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func parseASTFile(filePath string) *ast.File {
	buf := p.FileReader.Read(filePath)
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, filePath, buf, 0)
	if err != nil {
		panic(err)
	}
	return astFile
}

func parseASTStructs(file *ast.File) (typeNameAndStruct map[string]*ast.StructType) {
	typeNameAndStruct = map[string]*ast.StructType{}
	ast.Inspect(file, func(node ast.Node) bool {
		lastChildNode := node == nil
		if lastChildNode {
			return false
		}

		typeSpec, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		name := typeSpec.Name.Name
		typeNameAndStruct[name] = structType
		return true
	})
	return
}
