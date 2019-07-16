package reader

import (
	"testing"
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
		{
			name: "User to Users",
			args: args{
				str: "User",
			},
			want: "Users",
		},
		{
			name: "CoordinateItem to CoordinateItems",
			args: args{
				str: "CoordinateItem",
			},
			want: "CoordinateItems",
		},
		{
			name: "XyzJSON to XyzJsons",
			args: args{
				str: "XyzJSON",
			},
			want: "XyzJsons",
		},
		{
			name: "XJSON to XJsons",
			args: args{
				str: "XJSON",
			},
			want: "XJsons",
		},
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
