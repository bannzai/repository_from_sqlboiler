package strutil

import (
	"testing"
)

func TestSnakeCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "lcc to snake",
			args: args{
				str: "userRepository",
			},
			want: "user_repository",
		},
		{
			name: "ucc to snake",
			args: args{
				str: "UserRepository",
			},
			want: "user_repository",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.args.str); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_camelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ucc to lcc",
			args: args{
				str: "UserRepository",
			},
			want: "userRepository",
		},
		{
			name: "snake to lcc",
			args: args{
				str: "user_repository",
			},
			want: "userRepository",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := camelCase(tt.args.str); got != tt.want {
				t.Errorf("camelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ucc to lcc",
			args: args{
				str: "UserRepository",
			},
			want: "userRepository",
		},
		{
			name: "snake to lcc",
			args: args{
				str: "user_repository",
			},
			want: "userRepository",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerCamelCase(tt.args.str); got != tt.want {
				t.Errorf("LowerCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamelCase(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "lcc to ucc",
			args: args{
				target: "userRepository",
			},
			want: "UserRepository",
		},
		{
			name: "snake to ucc",
			args: args{
				target: "user_repository",
			},
			want: "UserRepository",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperCamelCase(tt.args.target); got != tt.want {
				t.Errorf("UpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecializeLowerCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "URL to url",
			args: args{
				str: "URL",
			},
			want: "url",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpecializeLowerCamelCase(tt.args.str); got != tt.want {
				t.Errorf("SpecializeLowerCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecializeUpperCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "url to URL",
			args: args{
				str: "url",
			},
			want: "URL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpecializeUpperCamelCase(tt.args.str); got != tt.want {
				t.Errorf("SpecializeUpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_EscapedReservedWord(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Add _ for Reserved word",
			args: args{
				str: "type",
			},
			want: "_type",
		},
		{
			name: "Not convert keyword",
			args: args{
				str: "xx",
			},
			want: "xx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EscapedReservedWord(tt.args.str); got != tt.want {
				t.Errorf("EscapedReservedWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEscapedReservedWord(t *testing.T) {
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
			if got := EscapedReservedWord(tt.args.str); got != tt.want {
				t.Errorf("EscapedReservedWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlural(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "user to users",
			args: args{
				str: "user",
			},
			want: "users",
		},
		{
			name: "community to communities",
			args: args{
				str: "community",
			},
			want: "communities",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Plural(tt.args.str); got != tt.want {
				t.Errorf("Plural() = %v, want %v", got, tt.want)
			}
		})
	}
}
