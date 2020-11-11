package testenv

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSetEnvironmentVariableFromReader(t *testing.T) {
	assert.NotContains(t, os.Environ(), "SAMPLE=test.1")
	r := strings.NewReader("SAMPLE=test.1\n")
	defer PatchEnvFromReader(t, r)()

	assert.Contains(t, os.Environ(), "SAMPLE=test.1")
}
func TestShouldAcceptEmptyEnvKey(t *testing.T) {
	assert.NotContains(t, os.Environ(), "SAMPLE")
	r := strings.NewReader("SAMPLE=\n")
	defer PatchEnvFromReader(t, r)()

	assert.NotContains(t, os.Environ(), "SAMPLE")
}

func TestShouldSetEnvVarFromMap(t *testing.T) {
	assert.NotContains(t, os.Environ(), "SAMPLE=test.1")
	m := map[string]string{
		"SAMPLE": "test.1",
	}
	defer PatchEnv(t, m)()

	assert.Contains(t, os.Environ(), "SAMPLE=test.1")
}

func TestShouldAcceptEmptyMap(t *testing.T) {
	defer PatchEnv(t, map[string]string{})()
}
