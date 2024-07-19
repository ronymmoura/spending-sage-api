package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	randomInt := RandomInt(1, 1000)
	require.NotEmpty(t, randomInt)
	require.Positive(t, randomInt)
	require.GreaterOrEqual(t, randomInt, int64(1))
	require.LessOrEqual(t, randomInt, int64(1000))
}

func TestRandomString(t *testing.T) {
	randomString := RandomString(10)
	require.NotEmpty(t, randomString)
	require.Len(t, randomString, 10)
}

func TestRandomOwner(t *testing.T) {
	randomOwner := RandomOwner()
	require.NotEmpty(t, randomOwner)
	require.Len(t, randomOwner, 6)
}

func TestRandomMoney(t *testing.T) {
	randomMoney := RandomMoney()
	require.NotEmpty(t, randomMoney)
	require.Positive(t, randomMoney)
	require.GreaterOrEqual(t, randomMoney, float64(1))
	require.LessOrEqual(t, randomMoney, float64(1000))
}

func TestRandomEmail(t *testing.T) {
	randomEmail := RandomEmail()
	require.NotEmpty(t, randomEmail)
}
