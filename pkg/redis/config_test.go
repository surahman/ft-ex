package redis

import (
	"errors"
	"strconv"
	"testing"

	"github.com/rs/xid"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"github.com/surahman/FTeX/pkg/constants"
	"github.com/surahman/FTeX/pkg/validator"
)

func TestRedisConfigs_Load(t *testing.T) {
	envAuthKey := constants.RedisPrefix() + "_AUTHENTICATION."
	envConnKey := constants.RedisPrefix() + "_CONNECTION."

	testCases := []struct {
		name         string
		input        string
		expectErrCnt int
		expectErr    require.ErrorAssertionFunc
	}{
		{
			name:         "empty - etc dir",
			input:        redisConfigTestData["empty"],
			expectErrCnt: 6,
			expectErr:    require.Error,
		}, {
			name:         "valid - etc dir",
			input:        redisConfigTestData["valid"],
			expectErrCnt: 0,
			expectErr:    require.NoError,
		}, {
			name:         "github-ci-runner - etc dir",
			input:        redisConfigTestData["github-ci-runner"],
			expectErrCnt: 0,
			expectErr:    require.NoError,
		}, {
			name:         "no username - etc dir",
			input:        redisConfigTestData["username_empty"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "no password - etc dir",
			input:        redisConfigTestData["password_empty"],
			expectErrCnt: 0,
			expectErr:    require.NoError,
		}, {
			name:         "no addrs - etc dir",
			input:        redisConfigTestData["no_addr"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "invalid max retries - etc dir",
			input:        redisConfigTestData["invalid_max_retries"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "invalid pool size - etc dir",
			input:        redisConfigTestData["invalid_pool_size"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "invalid min idle conns - etc dir",
			input:        redisConfigTestData["invalid_min_idle_conns"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "no max idle conns - etc dir",
			input:        redisConfigTestData["no_max_idle_conns"],
			expectErrCnt: 0,
			expectErr:    require.NoError,
		},
	}

	for _, testCase := range testCases {
		test := testCase
		t.Run(test.name, func(t *testing.T) {
			// Configure mock filesystem.
			fs := afero.NewMemMapFs()

			require.NoError(t, fs.MkdirAll(constants.EtcDir(), 0644), "Failed to create in memory directory")
			require.NoError(t, afero.WriteFile(fs, constants.EtcDir()+constants.RedisFileName(),
				[]byte(test.input), 0644), "Failed to write in memory file")

			// Load from mock filesystem.
			actual := &config{}
			err := actual.Load(fs)
			test.expectErr(t, err)

			validationError := &validator.ValidationError{}
			if errors.As(err, &validationError) {
				require.Lenf(t, validationError.Errors, test.expectErrCnt,
					"expected errors count is incorrect: %v", err)

				return
			}

			// Test configuring of environment variable.
			username := xid.New().String()
			password := xid.New().String()

			t.Setenv(envAuthKey+"USERNAME", username)
			t.Setenv(envAuthKey+"PASSWORD", password)

			addr := xid.New().String()
			maxConnAttempts := 12
			maxRetries := 55
			poolSize := 164
			minIdleConns := 9
			maxIdleConns := 101

			t.Setenv(envConnKey+"ADDR", addr)
			t.Setenv(envConnKey+"MAXCONNATTEMPTS", strconv.Itoa(maxConnAttempts))
			t.Setenv(envConnKey+"MAXRETRIES", strconv.Itoa(maxRetries))
			t.Setenv(envConnKey+"POOLSIZE", strconv.Itoa(poolSize))
			t.Setenv(envConnKey+"MINIDLECONNS", strconv.Itoa(minIdleConns))
			t.Setenv(envConnKey+"MAXIDLECONNS", strconv.Itoa(maxIdleConns))

			err = actual.Load(fs)
			require.NoErrorf(t, actual.Load(fs), "failed to load configurations file: %v", err)

			require.Equal(t, username, actual.Authentication.Username, "failed to load username.")
			require.Equal(t, password, actual.Authentication.Password, "failed to load password.")

			require.Equal(t, addr, actual.Connection.Addr, "failed to load address.")
			require.Equal(t, maxRetries, actual.Connection.MaxRetries, "failed to load max retries.")
			require.Equal(t, maxConnAttempts, actual.Connection.MaxConnAttempts,
				"failed to load max connection attempts.")
			require.Equal(t, poolSize, actual.Connection.PoolSize, "failed to load pool size.")
			require.Equal(t, minIdleConns, actual.Connection.MinIdleConns, "failed to load min idle conns.")
			require.Equal(t, maxIdleConns, actual.Connection.MaxIdleConns, "failed to load max idle conns.")
		})
	}
}
