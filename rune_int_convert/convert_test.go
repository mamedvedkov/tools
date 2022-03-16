package rune_int_convert

import "testing"

func TestConvertRune(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A",
			args: args{r: 'A'},
			want: 1,
		},
		{
			name: "B",
			args: args{r: 'B'},
			want: 2,
		},
		{
			name: "C",
			args: args{r: 'C'},
			want: 3,
		},
		{
			name: "Z",
			args: args{r: 'Z'},
			want: 26,
		},
		{
			name: "a",
			args: args{r: 'a'},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertRune(tt.args.r); got != tt.want {
				t.Errorf("ConvertRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Aa",
			args: args{"Aa"},
			want: -1,
		},
		{
			name: "AA",
			args: args{"AA"},
			want: 27,
		},
		{
			name: "AB",
			args: args{"AB"},
			want: 28,
		},
		{
			name: "BA",
			args: args{"BA"},
			want: 53,
		},
		{
			name: "AAA",
			args: args{"AAA"},
			want: 703,
		},
		{
			name: "AAB",
			args: args{"AAB"},
			want: 704,
		},
		{
			name: "A",
			args: args{"A"},
			want: 1,
		},
		{
			name: "B",
			args: args{"B"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertString(tt.args.s); got != tt.want {
				t.Errorf("ConvertString() = %v, want %v", got, tt.want)
			}
		})
	}
}
