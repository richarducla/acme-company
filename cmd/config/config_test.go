package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Config(t *testing.T) {
	t.Run("Defaults", func(t *testing.T) {
		c, err := FromEnv()
		assert.NoError(t, err)

		assert.Equal(t, "0.0.0.0", c.SrvHost)
		assert.Equal(t, "8080", c.SrvPort)
		assert.Equal(t, "dev", c.Env)
	})

	t.Run("SrvAddr", func(t *testing.T) {
		c := Config{
			SrvHost: "localhost",
			SrvPort: "8080",
		}
		assert.Equal(t, "localhost:8080", c.SrvAddr())
	})
}
