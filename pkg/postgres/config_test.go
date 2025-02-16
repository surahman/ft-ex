package postgres

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/rs/xid"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"github.com/surahman/FTeX/pkg/constants"
	"github.com/surahman/FTeX/pkg/validator"
	"gopkg.in/yaml.v3"
)

func TestNewConfig(t *testing.T) {
	cfg := newConfig()
	require.Equal(t, reflect.TypeOf(config{}), reflect.TypeOf(cfg), "new config type mismatch")
}

func TestConfigLoader(t *testing.T) {
	envAuthKey := constants.PostgresPrefix() + "_AUTHENTICATION."
	envConnKey := constants.PostgresPrefix() + "_CONNECTION."
	envPoolKey := constants.PostgresPrefix() + "_POOL."

	testCases := []struct {
		name      string
		input     string
		expectErr require.ErrorAssertionFunc
		expectLen int
	}{
		{
			name:      "empty - etc dir",
			input:     postgresConfigTestData["empty"],
			expectErr: require.Error,
			expectLen: 9,
		}, {
			name:      "valid - etc dir",
			input:     postgresConfigTestData["valid"],
			expectErr: require.NoError,
			expectLen: 0,
		}, {
			name:      "bad health check",
			input:     postgresConfigTestData["bad_health_check"],
			expectErr: require.Error,
			expectLen: 1,
		}, {
			name:      "invalid connections",
			input:     postgresConfigTestData["invalid_conns"],
			expectErr: require.Error,
			expectLen: 2,
		}, {
			name:      "invalid max connection attempts",
			input:     postgresConfigTestData["invalid_max_conn_attempts"],
			expectErr: require.Error,
			expectLen: 1,
		}, {
			name:      "invalid timeout",
			input:     postgresConfigTestData["invalid_timeout"],
			expectErr: require.Error,
			expectLen: 1,
		},
	}

	for _, testCase := range testCases {
		test := testCase
		t.Run(test.name, func(t *testing.T) {
			// Configure mock filesystem.
			fs := afero.NewMemMapFs()
			require.NoError(t, fs.MkdirAll(constants.EtcDir(), 0644), "Failed to create in memory directory")
			require.NoError(t, afero.WriteFile(fs, constants.EtcDir()+constants.PostgresFileName(),
				[]byte(test.input), 0644), "Failed to write in memory file")

			// Load from mock filesystem.
			actual := &config{}
			err := actual.Load(fs)
			test.expectErr(t, err)

			validationError := &validator.ValidationError{}
			if errors.As(err, &validationError) {
				require.Lenf(t, validationError.Errors, test.expectLen,
					"Expected errors count is incorrect: %v", err)

				return
			}

			// Load expected struct.
			expected := &config{}
			require.NoError(t, yaml.Unmarshal([]byte(test.input), expected),
				"failed to unmarshal expected constants")
			require.Truef(t, reflect.DeepEqual(expected, actual),
				"configurations loaded from disk do not match, expected %v, actual %v", expected, actual)

			// Test configuring of environment variable.
			username := xid.New().String()
			password := xid.New().String()

			t.Setenv(envAuthKey+"USERNAME", username)
			t.Setenv(envAuthKey+"PASSWORD", password)

			database := xid.New().String()
			host := xid.New().String()
			port := 5555
			timeout := 47
			maxConnAttempts := 9

			t.Setenv(envConnKey+"DATABASE", database)
			t.Setenv(envConnKey+"HOST", host)
			t.Setenv(envConnKey+"MAXCONNECTIONATTEMPTS", strconv.Itoa(maxConnAttempts))
			t.Setenv(envConnKey+"PORT", strconv.Itoa(port))
			t.Setenv(envConnKey+"TIMEOUT", strconv.Itoa(timeout))

			healthCheckPeriod := 13 * time.Second
			maxConns := 60
			minConns := 40

			t.Setenv(envPoolKey+"HEALTHCHECKPERIOD", healthCheckPeriod.String())
			t.Setenv(envPoolKey+"MAXCONNS", strconv.Itoa(maxConns))
			t.Setenv(envPoolKey+"MINCONNS", strconv.Itoa(minConns))

			require.NoErrorf(t, actual.Load(fs), "Failed to load constants file: %v", err)

			require.Equal(t, username, actual.Authentication.Username, "failed to load username")
			require.Equal(t, password, actual.Authentication.Password, "failed to load password")

			require.Equal(t, database, actual.Connection.Database, "failed to load database")
			require.Equal(t, host, actual.Connection.Host, "failed to load host")
			require.Equal(t, maxConnAttempts, actual.Connection.MaxConnAttempts, "failed to max connection attempts")
			require.Equal(t, uint16(port), actual.Connection.Port, "failed to load port")
			require.Equal(t, timeout, actual.Connection.Timeout, "failed to load timeout")

			require.Equal(t, healthCheckPeriod, actual.Pool.HealthCheckPeriod, "failed to load duration")
			require.Equal(t, int32(maxConns), actual.Pool.MaxConns, "failed to load max conns")
			require.Equal(t, int32(minConns), actual.Pool.MinConns, "failed to load min conns")
		})
	}
}
