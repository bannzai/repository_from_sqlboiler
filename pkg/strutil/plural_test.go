package strutil

import "testing"

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
		{
			name: "JSON to JSONs",
			args: args{
				str: "JSON",
			},
			want: "JSONs",
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
