package parser

import (
	"go/ast"
	"os"
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
	workingDirectory, _ := os.Getwd()
	type args struct {
		file *ast.File
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Get primary keys from testdata/user.go User",
			args: args{
				file: parseASTFile(workingDirectory + "/testdata/user.go"),
			},
			want: []string{"id", "full_name"},
		},
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
	workingDirectory, _ := os.Getwd()
	type args struct {
		file       *ast.File
		entityName string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Get fields from testdata/user.go User",
			args: args{
				file:       parseASTFile(workingDirectory + "/testdata/user.go"),
				entityName: "User",
			},
			want: map[string]string{
				"ID":        "uint",
				"FullName":  "string",
				"CreatedAt": "time.Time",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseASTFieldAndType(tt.args.file, tt.args.entityName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseASTFieldAndType() = %v, want %v", got, tt.want)
			}
		})
	}
}
