package main

import "testing"

func Test_normalize(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{phone: "1234567890"}, want: "1234567890"},
		{name: "2", args: args{phone: "123 456 7891"}, want: "1234567891"},
		{name: "3", args: args{phone: "(123) 456 7892"}, want: "1234567892"},
		{name: "4", args: args{phone: "(123) 456-7893"}, want: "1234567893"},
		{name: "5", args: args{phone: "123-456-7894"}, want: "1234567894"},
		{name: "6", args: args{phone: "123-456-7890"}, want: "1234567890"},
		{name: "7", args: args{phone: "1234567892"}, want: "1234567892"},
		{name: "8", args: args{phone: "(123)456-7892"}, want: "1234567892"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize(tt.args.phone); got != tt.want {
				t.Errorf("normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
