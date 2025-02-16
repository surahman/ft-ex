package quotes

import (
	"errors"
	"testing"
	"time"

	"github.com/rs/xid"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"github.com/surahman/FTeX/pkg/constants"
	"github.com/surahman/FTeX/pkg/validator"
)

func TestQuotesConfigs_Load(t *testing.T) {
	envFiatKey := constants.QuotesPrefix() + "_FIATCURRENCY."
	envCryptoKey := constants.QuotesPrefix() + "_CRYPTOCURRENCY."
	envConnKey := constants.QuotesPrefix() + "_CONNECTION."

	testCases := []struct {
		name         string
		input        string
		envValue     string
		expectErrCnt int
		expectErr    require.ErrorAssertionFunc
	}{
		{
			name:         "empty - etc dir",
			input:        quotesConfigTestData["empty"],
			expectErrCnt: 8,
			expectErr:    require.Error,
		}, {
			name:         "valid - etc dir",
			input:        quotesConfigTestData["valid"],
			expectErrCnt: 0,
			expectErr:    require.NoError,
		}, {
			name:         "no api key fiat",
			input:        quotesConfigTestData["no fiat api key"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "no api header fiat",
			input:        quotesConfigTestData["no fiat header key"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "no api endpoint fiat",
			input:        quotesConfigTestData["no fiat api endpoint"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "no fiat",
			input:        quotesConfigTestData["no fiat"],
			expectErrCnt: 3,
			expectErr:    require.Error,
		}, {
			name:         "no api key crypto",
			input:        quotesConfigTestData["no crypto api key"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "no api header crypto",
			input:        quotesConfigTestData["no crypto header key"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "no api endpoint crypto",
			input:        quotesConfigTestData["no crypto api endpoint"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "no crypto",
			input:        quotesConfigTestData["no crypto"],
			expectErrCnt: 3,
			expectErr:    require.Error,
		}, {
			name:         "no connection user-agent",
			input:        quotesConfigTestData["no connection user-agent"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "no connection timeout",
			input:        quotesConfigTestData["no connection timeout"],
			expectErrCnt: 1,
			expectErr:    require.Error,
		}, {
			name:         "no connection",
			input:        quotesConfigTestData["no connection"],
			expectErrCnt: 2,
			expectErr:    require.Error,
		},
	}
	for _, testCase := range testCases {
		test := testCase
		t.Run(test.name, func(t *testing.T) {
			// Configure mock filesystem.
			fs := afero.NewMemMapFs()
			require.NoError(t, fs.MkdirAll(constants.EtcDir(), 0644), "Failed to create in memory directory")
			require.NoError(t, afero.WriteFile(fs, constants.EtcDir()+constants.QuotesFileName(),
				[]byte(test.input), 0644), "Failed to write in memory file")

			// Load from mock filesystem.
			actual := &config{}
			err := actual.Load(fs)
			test.expectErr(t, err, "error expectation failed after loading from mock filesystem.")

			validationError := &validator.ValidationError{}
			if errors.As(err, &validationError) {
				require.Lenf(t, validationError.Errors, test.expectErrCnt,
					"expected errors count is incorrect: %v", err)

				return
			}

			// Test configuring of environment variable.
			apiKeyFiat := xid.New().String()
			headerKeyFiat := xid.New().String()
			apiEndpointFiat := xid.New().String()

			t.Setenv(envFiatKey+"APIKEY", apiKeyFiat)
			t.Setenv(envFiatKey+"HEADERKEY", headerKeyFiat)
			t.Setenv(envFiatKey+"ENDPOINT", apiEndpointFiat)

			apiKeyCrypto := xid.New().String()
			headerKeyCrypto := xid.New().String()
			apiEndpointCrypto := xid.New().String()

			t.Setenv(envCryptoKey+"APIKEY", apiKeyCrypto)
			t.Setenv(envCryptoKey+"HEADERKEY", headerKeyCrypto)
			t.Setenv(envCryptoKey+"ENDPOINT", apiEndpointCrypto)

			timeout := 999 * time.Second
			userAgent := xid.New().String()

			t.Setenv(envConnKey+"TIMEOUT", timeout.String())
			t.Setenv(envConnKey+"USERAGENT", userAgent)

			require.NoErrorf(t, actual.Load(fs), "failed to load configurations file: %v", err)

			require.Equal(t, apiKeyFiat, actual.FiatCurrency.APIKey, "failed to load fiat API Key.")
			require.Equal(t, headerKeyFiat, actual.FiatCurrency.HeaderKey, "failed to load fiat Header Key.")
			require.Equal(t, apiEndpointFiat, actual.FiatCurrency.Endpoint, "failed to load fiat API endpoint.")

			require.Equal(t, apiKeyCrypto, actual.CryptoCurrency.APIKey, "failed to load crypto API Key.")
			require.Equal(t, headerKeyCrypto, actual.CryptoCurrency.HeaderKey, "failed to load crypto Header Key.")
			require.Equal(t, apiEndpointCrypto, actual.CryptoCurrency.Endpoint, "failed to load crypto API endpoint.")

			require.Equal(t, timeout, actual.Connection.Timeout, "failed to load timeout.")
			require.Equal(t, userAgent, actual.Connection.UserAgent, "failed to load user-agent.")
		})
	}
}
