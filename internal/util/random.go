package util

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomFloat(min, max float64) float64 {
	return math.Round((min + rand.Float64()*(max-min)*100)) / 100
}

func RandomString(size int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < size; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() float64 {
	return RandomFloat(1, 1000)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@%s.%s", RandomString(10), RandomString(5), "com")
}
