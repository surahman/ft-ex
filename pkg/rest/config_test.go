package rest

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/rs/xid"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"github.com/surahman/FTeX/pkg/constants"
	"github.com/surahman/FTeX/pkg/validator"
	"gopkg.in/yaml.v3"
)

func TestRestConfigs_Load(t *testing.T) {
	keyspaceGen := constants.GetHTTPRESTPrefix() + "_SERVER."
	keyspaceAuth := constants.GetHTTPRESTPrefix() + "_AUTHORIZATION."

	testCases := []struct {
		name         string
		input        string
		expectErr    require.ErrorAssertionFunc
		expectErrCnt int
	}{
		{
			name:         "empty - etc dir",
			input:        restConfigTestData["empty"],
			expectErr:    require.Error,
			expectErrCnt: 5,
		}, {
			name:         "valid - etc dir",
			input:        restConfigTestData["valid"],
			expectErr:    require.NoError,
			expectErrCnt: 0,
		}, {
			name:         "out of range port - etc dir",
			input:        restConfigTestData["out of range port"],
			expectErr:    require.Error,
			expectErrCnt: 1,
		}, {
			name:         "no base path - etc dir",
			input:        restConfigTestData["no base path"],
			expectErr:    require.Error,
			expectErrCnt: 1,
		}, {
			name:         "no swagger path - etc dir",
			input:        restConfigTestData["no swagger path"],
			expectErr:    require.Error,
			expectErrCnt: 1,
		}, {
			name:         "no auth header - etc dir",
			input:        restConfigTestData["no auth header"],
			expectErr:    require.Error,
			expectErrCnt: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Configure mock filesystem.
			fs := afero.NewMemMapFs()
			require.NoError(t, fs.MkdirAll(constants.GetEtcDir(), 0644), "Failed to create in memory directory")
			require.NoError(t, afero.WriteFile(fs, constants.GetEtcDir()+constants.GetHTTPRESTFileName(),
				[]byte(testCase.input), 0644), "Failed to write in memory file")

			// Load from mock filesystem.
			actual := &config{}
			err := actual.Load(fs)
			testCase.expectErr(t, err)

			validationError := &validator.ValidationError{}
			if errors.As(err, &validationError) {
				require.Equalf(t, testCase.expectErrCnt, len(validationError.Errors),
					"expected errors count is incorrect: %v", err)

				return
			}

			// Load expected struct.
			expected := &config{}
			require.NoError(t, yaml.Unmarshal([]byte(testCase.input), expected), "failed to unmarshal expected constants")
			require.True(t, reflect.DeepEqual(expected, actual))

			// Test configuring of environment variable.
			basePath := xid.New().String()
			swaggerPath := xid.New().String()
			headerKey := xid.New().String()
			portNumber := 1600
			shutdownDelay := 36
			t.Setenv(keyspaceGen+"BASEPATH", basePath)
			t.Setenv(keyspaceGen+"SWAGGERPATH", swaggerPath)
			t.Setenv(keyspaceGen+"PORTNUMBER", strconv.Itoa(portNumber))
			t.Setenv(keyspaceGen+"SHUTDOWNDELAY", strconv.Itoa(shutdownDelay))
			t.Setenv(keyspaceAuth+"HEADERKEY", headerKey)
			err = actual.Load(fs)
			require.NoErrorf(t, err, "Failed to load constants file: %v", err)
			require.Equal(t, basePath, actual.Server.BasePath, "Failed to load base path environment variable into configs")
			require.Equal(t, swaggerPath, actual.Server.SwaggerPath,
				"Failed to load swagger path environment variable into configs")
			require.Equal(t, portNumber, actual.Server.PortNumber, "Failed to load port environment variable into configs")
			require.Equal(t, shutdownDelay, actual.Server.ShutdownDelay,
				"Failed to load shutdown delay environment variable into configs")
			require.Equal(t, headerKey, actual.Authorization.HeaderKey,
				"Failed to load authorization header key environment variable into configs")
		})
	}
}
