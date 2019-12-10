package codegen

import (
	crand "crypto/rand"
	"encoding/binary"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"math"
	"math/rand"
	"strings"
)

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		logger.Error(err)
	}
	return v
}

//return random int in the range min...max
func randomInt(min, max int) int {
	var src cryptoSource
	rnd := rand.New(src)
	if rnd == nil {
		return min + rand.Intn(max-min)
	}
	return min + rnd.Intn(max-min)
}

//return random char string from charset
func randomChar(cs []byte) string {
	return string(cs[randomInt(0, len(cs)-1)])
}

//repeat string with one str (#)
func repeatStr(count int, str string) string {
	return strings.Repeat(str, count)
}

func isFeasible(charset CharsetType, pattern string, count int) bool {
	ls := strings.Count(pattern, "#")
	if math.Pow(float64(len(charset)), float64(ls)) >= float64(count) {
		return true
	}
	return false
}
