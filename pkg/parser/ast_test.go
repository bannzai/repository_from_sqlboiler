package parser

import (
	"go/ast"
	"reflect"
	"testing"
)

func Test_parseASTFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want *ast.File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseASTFile(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseASTFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseASTPrimaryKeyFields(t *testing.T) {
	type args struct {
		file *ast.File
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseASTPrimaryKeyFields(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseASTPrimaryKeyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseASTFieldAndType(t *testing.T) {
	type args struct {
		file       *ast.File
		entityName string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseASTFieldAndType(tt.args.file, tt.args.entityName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseASTFieldAndType() = %v, want %v", got, tt.want)
			}
		})
	}
}
