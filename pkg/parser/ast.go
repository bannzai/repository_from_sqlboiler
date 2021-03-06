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

func parseASTPrimaryKeyFields(file *ast.File) []string {
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
				uniqueKeys = append(uniqueKeys, strings.ReplaceAll(basicLit.Value, "\"", ""))
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
			if fieldName == "R" {
				continue
			}
			if fieldName == "L" {
				continue
			}
			switch fieldType := f.Type.(type) {
			case *ast.Ident:
				fieldAndType[fieldName] = fieldType.Name
			case *ast.SelectorExpr:
				packageName, ok := fieldType.X.(*ast.Ident)
				if !ok {
					fmt.Printf("[WARNING]⚠  Unexpected field type : %v\n", f.Type)
					continue
				}
				selector := fieldType.Sel
				fieldAndType[fieldName] = packageName.Name + "." + selector.Name
			default:
				fmt.Printf("[WARNING]⚠  Unexpected field type : %v\n", f.Type)
				continue
			}
		}

		return continueTraverse
	})

	return fieldAndType
}

func parseASTBaseStructName(file *ast.File) string {
	name := ""
	ast.Inspect(file, func(node ast.Node) bool {
		if len(name) > 0 {
			return endTraverse
		}
		lastChildNode := node == nil
		if lastChildNode {
			return endTraverse
		}

		genDecl, ok := node.(*ast.GenDecl)
		if !ok {
			return continueTraverse
		}

		typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
		if !ok {
			return continueTraverse
		}

		// FIXME: Get first typeSpec.Name.Name. But want other good idea.
		name = typeSpec.Name.Name
		return continueTraverse
	})

	return name
}
