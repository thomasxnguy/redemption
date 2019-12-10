package codegen

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"math/rand"
	"strings"
	"time"
)

type CharsetType string

const (
	CharsetNumbers      CharsetType = "0123456789"
	CharsetAlphabetic   CharsetType = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetAlphanumeric CharsetType = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// defaults
	minLength int = 6
)

//initialize package base config
func init() {
	rand.Seed(time.Now().UnixNano())
}

type Options struct {
	Charset CharsetType
	Prefix  string
	Pattern string
}

type Generator struct {
	Count   int
	Options *Options
}

func GenerateCodes(count int, options *Options) ([]string, error) {
	gen := newGenerator(count, options)
	result, err := gen.Run()
	if err != nil {
		return nil, errors.E(err, "generate codes failed")
	}
	return *result, nil
}

//New generator config
func newGenerator(count int, options *Options) (gen *Generator) {
	// check pattern
	if len(options.Pattern) == 0 {
		options.Pattern = repeatStr(minLength, "#")
	} else if len(options.Pattern) < minLength {
		options.Pattern += repeatStr(minLength-len(options.Pattern), "#")
	}

	// check charset
	if options.Charset == CharsetType("") {
		options.Charset = CharsetAlphanumeric
	}

	return &Generator{
		Count:   count,
		Options: options,
	}
}

//Run voucher code generator
func (g *Generator) Run() (*[]string, error) {
	if !isFeasible(g.Options.Charset, g.Options.Pattern, g.Count) {
		return nil, errors.E("Not possible to generate requested number of codes")
	}
	result := make([]string, g.Count)
	mapping := make(map[string]bool)
	for i := 0; i < g.Count; i++ {
		code := g.one()
		_, ok := mapping[code]
		if ok {
			return nil, errors.E("failed to generate all codes do you have a code collision")
		}
		mapping[code] = true
		result[i] = code
	}
	return &result, nil
}

// generate one vouchers code
func (g *Generator) one() string {
	pts := strings.Split(g.Options.Pattern, "")
	for i, v := range pts {
		if v == "#" {
			pts[i] = randomChar([]byte(g.Options.Charset))
		}
	}
	return g.Options.Prefix + strings.Join(pts, "")
}
