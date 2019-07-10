package strutil

import "testing"

func TestSnakeCase(t *testing.T) {
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperCamelCase(tt.args.target); got != tt.want {
				t.Errorf("UpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_specializeIfMatched(t *testing.T) {
	type args struct {
		str                string
		specliazedKeywords []string
		specialize         func(string) string
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
			if got := specializeIfMatched(tt.args.str, tt.args.specliazedKeywords, tt.args.specialize); got != tt.want {
				t.Errorf("specializeIfMatched() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpecializeUpperCamelCase(tt.args.str); got != tt.want {
				t.Errorf("SpecializeUpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
