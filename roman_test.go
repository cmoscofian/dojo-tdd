package main

import "testing"

func TestConvertToRoman(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		args int
		want string
	}{
		{
			args: 1,
			want: "I",
		},
		{
			args: 2,
			want: "II",
		},
		{
			args: 3,
			want: "III",
		},
		{
			args: 4,
			want: "IV",
		},
		{
			args: 5,
			want: "V",
		},
		{
			args: 6,
			want: "VI",
		},
		{
			args: 8,
			want: "VIII",
		},
		{
			args: 9,
			want: "IX",
		},
		{
			args: 10,
			want: "X",
		},
		{
			args: 11,
			want: "XI",
		},
		{
			args: 13,
			want: "XIII",
		},
		{
			args: 14,
			want: "XIV",
		},
		{
			args: 19,
			want: "XIX",
		},
		{
			args: 20,
			want: "XX",
		},
		{
			args: 40,
			want: "XL",
		},
		{
			args: 39,
			want: "XXXIX",
		},
		{
			args: 50,
			want: "L",
		},
		{
			args: 88,
			want: "LXXXVIII",
		},
		{
			args: 90,
			want: "XC",
		},
		{
			args: 246,
			want: "CCXLVI",
		},
		{
			args: 2421,
			want: "MMCDXXI",
		},
		{
			args: 3999,
			want: "MMMCMXCIX",
		},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := ConvertToRoman(tt.args); got != tt.want {
				t.Errorf("ConvertToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
