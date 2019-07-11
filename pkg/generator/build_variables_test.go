package generator

import (
	"testing"

	"github.com/bannzai/repository_from_sqlboiler/pkg/model"
)

func Test_fetchByPrimaryKey(t *testing.T) {
	type args struct {
		entity model.Entity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
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
	type args struct {
		entity model.Entity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
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
	type args struct {
		entity model.Entity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
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
	type args struct {
		entity model.Entity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
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
	type args struct {
		entity model.Entity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sqlQueryArguments(tt.args.entity); got != tt.want {
				t.Errorf("sqlQueryArguments() = %v, want %v", got, tt.want)
			}
		})
	}
}
