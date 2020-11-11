package testenv

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

// PatchEnvFromFile uses a .env file to patch the environment
func PatchEnvFromFile(t *testing.T, envfile string) func() {
	_, err := os.Stat(envfile)
	if err == nil || !os.IsNotExist(err) {
		// if stat returns an error, but it does NOT say the file does not exists,
		// invoke godotenv.Load and let it return the error
		patchEnv, err := godotenv.Read(envfile)
		if err != nil {
			t.Error(err)
		}
		return PatchEnv(t, patchEnv)
	}

	return PatchEnv(t, nil)
}

func PatchEnvFromReader(t *testing.T, reader io.Reader) func() {
	patchEnv, err := godotenv.Parse(reader)
	if err != nil {
		t.Error(err)
	}
	return PatchEnv(t, patchEnv)
}

// PatchEnv will patch the environment using the given map
func PatchEnv(t *testing.T, envMap map[string]string) func() {
	if len(envMap) == 0 {
		return func() {}
	}
	oldEnv := os.Environ()

	setEnv(t, envMap)

	return func() {
		os.Clearenv()
		setEnv(t, ToMap(oldEnv))
	}
}

func setEnv(t *testing.T, env map[string]string) {
	for key, value := range env {
		if err := os.Setenv(key, value); err != nil {
			t.Error(err)
		}
	}
}

func ToMap(env []string) map[string]string {
	m := map[string]string{}
	for _, kv := range env {
		k, v := splitEnvKeyValue(kv)
		m[k] = v
	}

	return m
}

func splitEnvKeyValue(kv string) (string, string) {
	switch {
	case kv == "":
		return "", ""
	case strings.HasPrefix(kv, "="):
		k, v := splitEnvKeyValue(kv[1:])
		return "=" + k, v
	case strings.Contains(kv, "="):
		parts := strings.SplitN(kv, "=", 2)
		return parts[0], parts[1]
	default:
		return kv, ""
	}
}
