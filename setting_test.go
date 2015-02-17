package setting

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SimpleConfig struct {
	AAA string `default:"A"`
	BBB string `env:"BBB"`
	CCC string `env:"CCC" default:"C"`
}

func TestSettingFromDefault(t *testing.T) {
	config := new(SimpleConfig)
	Load(config)

	assert.Equal(t, config.AAA, "A")
}

func TestSettingFromEnv(t *testing.T) {
	os.Setenv("BBB", "three b's")
	defer os.Unsetenv("BBB")

	config := new(SimpleConfig)
	Load(config)

	assert.Equal(t, config.BBB, "three b's")
}

func TestSettingFromEnvDefaultFallback(t *testing.T) {
	config := new(SimpleConfig)
	Load(config)

	assert.Equal(t, config.CCC, "C")
}
