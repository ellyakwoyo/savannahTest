package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigLoader(t *testing.T) {
	err := ConfigLoader()
	assert.NoError(t, err, "Error loading config")
}
