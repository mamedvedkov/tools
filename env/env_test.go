package env

import (
	"os"
	"testing"

	"github.com/go-logr/zapr"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var logger = zapr.NewLogger(zap.NewExample())

func TestEnv_String_found(t *testing.T) {
	required := require.New(t)

	key := "TEST_ENV"
	expectedValue := "test_val"

	err := os.Setenv(key, expectedValue)
	required.NoError(err)

	defer os.Unsetenv(key)

	actualValue := GetWithOption(key, WithLogging(logger)).String("def")
	required.Equal(expectedValue, actualValue)
}

func TestEnv_String_not_found(t *testing.T) {
	required := require.New(t)

	key := "TEST_ENV"
	expectedValue := "def"
	actualValue := GetWithOption(key, WithLogging(logger)).String(expectedValue)
	required.Equal(expectedValue, actualValue)
}

func TestEnv_MustString_found(t *testing.T) {
	required := require.New(t)

	key := "TEST_ENV"
	expectedValue := "test_val"

	err := os.Setenv(key, expectedValue)
	required.NoError(err)

	defer os.Unsetenv(key)

	actualValue := GetWithOption(key, WithLogging(logger)).MustString()
	required.Equal(expectedValue, actualValue)
}

func TestEnv_MustString_not_found(t *testing.T) {
	required := require.New(t)

	key := "TEST_ENV"

	required.Panics(func() {
		GetWithOption(key, WithLogging(logger)).MustString()
	})
}

func TestEnv_Int_found(t *testing.T) {
	required := require.New(t)

	key := "TEST_ENV"
	expectedValue := "42"

	err := os.Setenv(key, expectedValue)
	required.NoError(err)

	defer os.Unsetenv(key)

	actualValue := GetWithOption(key, WithLogging(logger)).Int(42)
	required.Equal(42, actualValue)
}

func TestEnv_Int_not_found(t *testing.T) {
	required := require.New(t)

	key := "TEST_ENV"
	expectedValue := 42
	actualValue := GetWithOption(key, WithLogging(logger)).Int(expectedValue)
	required.Equal(expectedValue, actualValue)
}

func TestEnv_MustInt_found(t *testing.T) {
	required := require.New(t)

	key := "TEST_ENV"
	expectedValue := "42"

	err := os.Setenv(key, expectedValue)
	required.NoError(err)

	defer os.Unsetenv(key)

	actualValue := GetWithOption(key, WithLogging(logger)).MustInt()
	required.Equal(42, actualValue)
}

func TestEnv_MustInt_not_found(t *testing.T) {
	required := require.New(t)

	key := "TEST_ENV"

	required.Panics(func() {
		GetWithOption(key, WithLogging(logger)).MustInt()
	})
}
