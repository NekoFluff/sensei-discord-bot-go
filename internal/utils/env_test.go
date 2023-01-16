package utils

import (
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestUtils_GetEnvVar(t *testing.T) {
	os.Setenv("ENV_VAR", "^")
	actual := GetEnvVar("ENV_VAR")
	assert.Equal(t, "^", actual)
}

func TestUtils_GetEnvVar_Empty(t *testing.T) {
	actual := GetEnvVar("INVALID_ENV_VAR")
	assert.Equal(t, "", actual)
}
