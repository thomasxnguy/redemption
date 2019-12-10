package codegen

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerateCodes(t *testing.T) {
	type args struct {
		count   int
		options *Options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"numbers", args{20, &Options{Prefix: "numbers_", Charset: CharsetNumbers}}, false},
		{"custom charset", args{300, &Options{Prefix: "custom-char--", Charset: "123456789ABCDabfg"}}, false},
		{"custom small charset", args{50, &Options{Prefix: "custom-char--", Charset: "123Cbfg"}}, false},
		{"alphabetic", args{100, &Options{Prefix: "100-", Charset: CharsetAlphabetic}}, false},
		{"pattern", args{100, &Options{Prefix: "pattern-", Pattern: "###-###-###"}}, false},
		{"small-pattern", args{100, &Options{Prefix: "small-pattern-", Pattern: "##"}}, false},
		{"error code collision", args{50000, &Options{Prefix: "custom-char--", Charset: "12bfg"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateCodes(tt.args.count, tt.args.options)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.args.count, len(got))
			for _, g := range got {
				prefix := strings.HasPrefix(g, tt.args.options.Prefix)
				assert.True(t, prefix)
				length := len(tt.args.options.Prefix) + len(tt.args.options.Pattern)
				assert.Equal(t, length, len(g))
			}
		})
	}
}

func Test_newGenerator(t *testing.T) {
	type args struct {
		count   int
		options *Options
	}
	tests := []struct {
		name    string
		args    args
		wantGen *Generator
	}{
		{
			"test parse",
			args{10, &Options{
				Charset: CharsetAlphabetic,
				Prefix:  "trust-",
				Pattern: "###-####-###",
			}},
			&Generator{
				Count: 10,
				Options: &Options{
					Charset: CharsetAlphabetic,
					Prefix:  "trust-",
					Pattern: "###-####-###",
				},
			},
		}, {
			"test parse default",
			args{100, &Options{}},
			&Generator{
				Count: 100,
				Options: &Options{
					Charset: CharsetAlphanumeric,
					Prefix:  "",
					Pattern: "######",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGen := newGenerator(tt.args.count, tt.args.options)
			assert.Equal(t, tt.wantGen, gotGen)
		})
	}
}
