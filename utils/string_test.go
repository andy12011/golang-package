package utils

import (
	"testing"
)

func TestStringConnectByBuffer(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{args: []string{""}},
			want: "",
		},
		{
			name: "case2",
			args: args{args: []string{}},
			want: "",
		},
		{
			name: "case3",
			args: args{args: []string{"str1", "-", "str2", "-", "str3"}},
			want: "str1-str2-str3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringConnectByBuffer(tt.args.args...); got != tt.want {
				t.Errorf("StringConnectByBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringConnectByBuilder(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{args: []string{""}},
			want: "",
		},
		{
			name: "case2",
			args: args{args: []string{}},
			want: "",
		},
		{
			name: "case3",
			args: args{args: []string{"str1", "-", "str2", "-", "str3"}},
			want: "str1-str2-str3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringConnectByBuilder(tt.args.args...); got != tt.want {
				t.Errorf("StringConnectByBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
