package codegen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isFeasible(t *testing.T) {
	type args struct {
		charset CharsetType
		pattern string
		count   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Alphanumeric Charset", args{CharsetAlphanumeric, "#####", 10}, true},
		{"Alphabetic Charset", args{CharsetAlphabetic, "##-##", 100}, true},
		{"Numbers Charset", args{CharsetNumbers, "###-###-###", 1000}, true},
		{"Custom Charset", args{"abc123", "####", 3}, true},
		{"Empty Charset", args{"", "###-###-###", 3}, false},
		{"Empty Pattern", args{"abcde", "", 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFeasible(tt.args.charset, tt.args.pattern, tt.args.count); got != tt.want {
				t.Errorf("isFeasible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomChar(t *testing.T) {
	tests := []struct {
		name string
		cs   []byte
	}{
		{"random test 1", []byte("charset")},
		{"random test 2", []byte("truswallet")},
		{"random test 3", []byte("truswalletapp")},
		{"random test 4", []byte(CharsetAlphanumeric)},
		{"random test 5", []byte(CharsetAlphabetic)},
		{"random test 6", []byte(CharsetNumbers)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randomChar(tt.cs)
			assert.Len(t, got, 1)
		})
	}
}

func Test_randomInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{"0-100", args{0, 10}},
		{"10-100", args{10, 100}},
		{"0-1000", args{0, 100000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randomInt(tt.args.min, tt.args.max)
			assert.GreaterOrEqual(t, got, tt.args.min)
			assert.LessOrEqual(t, got, tt.args.max)
		})
	}
}

func Test_repeatStr(t *testing.T) {
	type args struct {
		count int
		str   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"4-#", args{4, "#"}, "####"},
		{"10-$", args{10, "$"}, "$$$$$$$$$$"},
		{"3-trust", args{3, "trust"}, "trusttrusttrust"},
		{"0-trust", args{0, "trust"}, ""},
		{"1000000-empty", args{1000000, ""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatStr(tt.args.count, tt.args.str); got != tt.want {
				t.Errorf("repeatStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
