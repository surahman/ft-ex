package logger

import (
	"errors"
	"reflect"
	"testing"

	"github.com/rs/xid"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"github.com/surahman/FTeX/pkg/constants"
	"github.com/surahman/FTeX/pkg/validator"
	"gopkg.in/yaml.v3"
)

func TestZapConfig_Load(t *testing.T) {
	envCfgKey := constants.LoggerPrefix() + "_BUILTINCONFIG"
	envEncKey := constants.LoggerPrefix() + "_BUILTINENCODERCONFIG"

	testCases := []struct {
		name      string
		input     string
		cfgKey    string
		encKey    string
		expectErr require.ErrorAssertionFunc
		expectNil require.ValueAssertionFunc
		expectLen int
	}{
		{
			name:      "invalid - empty",
			input:     loggerConfigTestData["empty"],
			cfgKey:    "Production",
			encKey:    "Production",
			expectErr: require.Error,
			expectNil: require.Nil,
			expectLen: 2,
		}, {
			name:      "invalid - builtin",
			input:     loggerConfigTestData["invalid_builtin"],
			cfgKey:    xid.New().String(),
			encKey:    xid.New().String(),
			expectErr: require.Error,
			expectNil: require.Nil,
			expectLen: 2,
		}, {
			name:      "valid - development",
			input:     loggerConfigTestData["valid_devel"],
			cfgKey:    "Production",
			encKey:    "Production",
			expectErr: require.NoError,
			expectNil: require.Nil,
			expectLen: 0,
		}, {
			name:      "valid - production",
			input:     loggerConfigTestData["valid_prod"],
			cfgKey:    "Development",
			encKey:    "Development",
			expectErr: require.NoError,
			expectNil: require.Nil,
			expectLen: 0,
		}, {
			name:      "valid - full constants",
			input:     loggerConfigTestData["valid_config"],
			cfgKey:    "Production",
			encKey:    "Production",
			expectErr: require.NoError,
			expectNil: require.NotNil,
			expectLen: 0,
		},
	}
	for _, testCase := range testCases {
		test := testCase
		t.Run(test.name, func(t *testing.T) {
			// Configure mock filesystem.
			fs := afero.NewMemMapFs()
			require.NoError(t, fs.MkdirAll(constants.EtcDir(), 0644), "Failed to create in memory directory")
			require.NoError(t, afero.WriteFile(fs, constants.EtcDir()+constants.LoggerFileName(),
				[]byte(test.input), 0644), "Failed to write in memory file")

			// Load from mock filesystem.
			actual := &config{}
			err := actual.Load(fs)
			test.expectErr(t, err)

			validationError := &validator.ValidationError{}
			if errors.As(err, &validationError) {
				require.Lenf(t, validationError.Errors, test.expectLen, "Expected errors count is incorrect: %v", err)

				return
			}

			// Load expected struct.
			expected := &config{}
			require.NoError(t, yaml.Unmarshal([]byte(test.input), expected), "failed to unmarshal expected constants")
			require.True(t, reflect.DeepEqual(expected, actual))

			// Test configuring of environment variable.
			t.Setenv(envCfgKey, test.cfgKey)
			t.Setenv(envEncKey, test.encKey)
			require.NoErrorf(t, actual.Load(fs), "Failed to load constants file: %v", err)

			require.Equalf(t, test.cfgKey, actual.BuiltinConfig, "Failed to load environment variable into constants")
			require.Equalf(t, test.encKey, actual.BuiltinEncoderConfig, "Failed to load environment variable into encoder")

			test.expectNil(t, actual.GeneralConfig, "Check for nil general constants failed")
		})
	}
}
