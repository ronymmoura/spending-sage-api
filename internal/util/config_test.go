package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCantLoadConfig(t *testing.T) {
	_, err := LoadConfig("a")
	require.Error(t, err)
}

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../../.env")

	require.NoError(t, err)
	require.NotEmpty(t, config)
	require.Equal(t, "development", config.Environment)
}
