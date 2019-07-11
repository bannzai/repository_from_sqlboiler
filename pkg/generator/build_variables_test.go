package generator

import (
	"os"
	"testing"

	"github.com/bannzai/repository_from_sqlboiler/pkg/model"
	"github.com/bannzai/repository_from_sqlboiler/pkg/parser"
)

func Test_fetchByPrimaryKey(t *testing.T) {
	workingDirectory, _ := os.Getwd()
	type args struct {
		entity model.Entity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Mapped primary key information",
			args: args{
				entity: parser.Entity{
					FilePath: workingDirectory + "/testdata/user.go",
				}.Parse(),
			},
			want: "FetchByIDAndTypeAndFullName(id uint, _type string, fullName string) entity.User",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fetchByPrimaryKey(tt.args.entity); got != tt.want {
				t.Errorf("fetchByPrimaryKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetchByPrimaryKeyFunctionName(t *testing.T) {
	workingDirectory, _ := os.Getwd()
	type args struct {
		entity model.Entity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Mapped primary key information",
			args: args{
				entity: parser.Entity{
					FilePath: workingDirectory + "/testdata/user.go",
				}.Parse(),
			},
			want: "FetchByIDAndTypeAndFullName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fetchByPrimaryKeyFunctionName(tt.args.entity); got != tt.want {
				t.Errorf("fetchByPrimaryKeyFunctionName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetchByPrimaryKeyFunctionArgument(t *testing.T) {
	workingDirectory, _ := os.Getwd()
	type args struct {
		entity model.Entity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Mapped primary key information",
			args: args{
				entity: parser.Entity{
					FilePath: workingDirectory + "/testdata/user.go",
				}.Parse(),
			},
			want: "id uint, _type string, fullName string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fetchByPrimaryKeyFunctionArgument(tt.args.entity); got != tt.want {
				t.Errorf("fetchByPrimaryKeyFunctionArgument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listOfPrimaryKeys(t *testing.T) {
	workingDirectory, _ := os.Getwd()
	type args struct {
		entity model.Entity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Mapped primary key information",
			args: args{
				entity: parser.Entity{
					FilePath: workingDirectory + "/testdata/user.go",
				}.Parse(),
			},
			want: "id, _type, full_name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listOfPrimaryKeys(tt.args.entity); got != tt.want {
				t.Errorf("listOfPrimaryKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlQueryArguments(t *testing.T) {
	workingDirectory, _ := os.Getwd()
	type args struct {
		entity model.Entity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Mapped primary key informations for where condition",
			args: args{
				entity: parser.Entity{
					FilePath: workingDirectory + "/testdata/user.go",
				}.Parse(),
			},
			want: "qm.Where(\"id=?\", id), qm.Where(\"type=?\", _type), qm.Where(\"full_name=?\", fullName)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sqlQueryArguments(tt.args.entity); got != tt.want {
				t.Errorf("sqlQueryArguments() = %v, want %v", got, tt.want)
			}
		})
	}
}
