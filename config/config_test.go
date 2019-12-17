package config_test

import (
	"testing"

	"github.com/donbattery/snake-hub-backend/config"
	"github.com/spf13/cobra"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	c := config.New()

	assert.NotNil(t, c, "Config object should be returned")

}

func TestInit(t *testing.T) {
	c := config.New()

	assert.NoError(t, c.Init(&cobra.Command{}, []string{}), "Init should not return any errors")

}
