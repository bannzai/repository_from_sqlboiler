package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

const (
	endTraverse      bool = false
	continueTraverse bool = true
)

func parseASTFile(filePath string) *ast.File {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("Can not read %v, and got error :%v", filePath, err))
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, filePath, data, 0)
	if err != nil {
		panic(err)
	}
	return astFile
}

func pearseASTPrimaryKeyFiels(file *ast.File) []string {
	uniqueKeys := []string{}

	ast.Inspect(file, func(node ast.Node) bool {
		lastChildNode := node == nil
		if lastChildNode {
			return endTraverse
		}

		valueSpec, ok := node.(*ast.ValueSpec)
		if !ok {
			return continueTraverse
		}

		if len(valueSpec.Names) != 1 {
			return endTraverse
		}

		if !strings.HasSuffix(valueSpec.Names[0].Name, "PrimaryKeyColumns") {
			return endTraverse
		}

		fieldsNamesExpr, ok := valueSpec.Values[0].(*ast.CompositeLit)
		if !ok {
			return continueTraverse
		}

		for _, element := range fieldsNamesExpr.Elts {
			if basicLit, ok := element.(*ast.BasicLit); ok {
				uniqueKeys = append(uniqueKeys, basicLit.Value)
			}
		}

		return continueTraverse
	})

	return uniqueKeys
}

func parseASTFieldAndType(file *ast.File, entityName string) map[string]string {
	fieldAndType := map[string]string{}
	ast.Inspect(file, func(node ast.Node) bool {
		lastChildNode := node == nil
		if lastChildNode {
			return endTraverse
		}

		typeSpec, ok := node.(*ast.TypeSpec)
		if !ok {
			return continueTraverse
		}

		if typeSpec.Name.Name != entityName {
			return continueTraverse
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return continueTraverse
		}

		for _, f := range structType.Fields.List {
			fieldName := f.Names[0].Name
			if fieldName == entityName+"R" {
				continue
			}
			if fieldName == entityName+"L" {
				continue
			}
			fieldType, ok := f.Type.(*ast.Ident)
			if !ok {
				fmt.Printf("[WARNING]⚠  Unexpected field type : %v\n", f.Type)
				continue
			}
			fieldAndType[fieldName] = fieldType.Name
		}

		return continueTraverse
	})

	return fieldAndType
}
