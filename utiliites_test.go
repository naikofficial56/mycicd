package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		num1 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test cases 1",
			args: args{num1: 24},
			want: 34,
		},
		{
			name: "test cases 2",
			args: args{num1: 24},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.num1); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
