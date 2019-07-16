package reader

import (
	"reflect"
	"testing"
	"text/template"
)

func Test_entitySelectorName(t *testing.T) {
	type args struct {
		str string
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
			if got := entitySelectorName(tt.args.str); got != tt.want {
				t.Errorf("entitySelectorName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_golangArgumentCase(t *testing.T) {
	type args struct {
		str string
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
			if got := golangArgumentCase(tt.args.str); got != tt.want {
				t.Errorf("golangArgumentCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_golangVariableCase(t *testing.T) {
	type args struct {
		str string
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
			if got := golangVariableCase(tt.args.str); got != tt.want {
				t.Errorf("golangVariableCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplate_Read(t *testing.T) {
	type fields struct {
		FilePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   *template.Template
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := Template{
				FilePath: tt.fields.FilePath,
			}
			if got := t.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
