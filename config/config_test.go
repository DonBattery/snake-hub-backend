package config_test

import (
	"testing"

	"github.com/donbattery/snake-hub-backend/config"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	c := config.New()

	assert.NotNil(t, c, "Config object should be returned")

}

func TestSetup(t *testing.T) {
	c := config.New()

	assert.NoError(t, c.Setup(nil, nil), "Setup should not return any errors")

}
