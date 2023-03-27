package postgres

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestFiat_CreateFiatAccount(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, clientID2 := resetTestFiatAccounts(t)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)

	defer cancel()

	// Account collisions test.
	for key, testCase := range getTestFiatAccounts(clientID1, clientID2) {
		parameters := testCase

		t.Run(fmt.Sprintf("Inserting %s", key), func(t *testing.T) {
			for _, param := range parameters {
				accInfo := param
				rowCount, err := connection.Query.createFiatAccount(ctx, &accInfo)
				require.Error(t, err, "did not error whilst inserting duplicate fiat account.")
				require.Equal(t, int64(0), rowCount, "rows were added.")
			}
		})
	}
}

func TestFiat_RowLockFiatAccount(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, clientID2 := resetTestFiatAccounts(t)

	// Get general ledger entry test cases.
	testCases := []struct {
		name         string
		parameter    rowLockFiatAccountParams
		errExpected  require.ErrorAssertionFunc
		boolExpected require.BoolAssertionFunc
	}{
		{
			name: "Client1 - USD",
			parameter: rowLockFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyUSD,
			},
			errExpected:  require.NoError,
			boolExpected: require.True,
		}, {
			name: "Client1 - AED",
			parameter: rowLockFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyAED,
			},
			errExpected:  require.NoError,
			boolExpected: require.True,
		}, {
			name: "Client1 - CAD",
			parameter: rowLockFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyCAD,
			},
			errExpected:  require.NoError,
			boolExpected: require.True,
		}, {
			name: "Client2 - USD",
			parameter: rowLockFiatAccountParams{
				ClientID: clientID2,
				Currency: CurrencyUSD,
			},
			errExpected:  require.NoError,
			boolExpected: require.True,
		}, {
			name: "Client2 - AED",
			parameter: rowLockFiatAccountParams{
				ClientID: clientID2,
				Currency: CurrencyAED,
			},
			errExpected:  require.NoError,
			boolExpected: require.True,
		}, {
			name: "Client2 - CAD",
			parameter: rowLockFiatAccountParams{
				ClientID: clientID2,
				Currency: CurrencyCAD,
			},
			errExpected:  require.NoError,
			boolExpected: require.True,
		}, {
			name: "Client1 - Not Found",
			parameter: rowLockFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyEUR,
			},
			errExpected:  require.Error,
			boolExpected: require.False,
		},
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)

	defer cancel()

	// Insert new fiat accounts.
	for _, testCase := range testCases {
		test := testCase

		t.Run(fmt.Sprintf("Inserting %s", test.name), func(t *testing.T) {
			balance, err := connection.Query.rowLockFiatAccount(ctx, &test.parameter)
			test.errExpected(t, err, "error expectation condition failed.")
			test.boolExpected(t, balance.Valid, "invalid balance.")
		})
	}
}

func TestFiat_UpdateBalanceFiatAccount(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, _ := resetTestFiatAccounts(t)

	// Test data setup.
	const expectedBalance = 4242.42

	var (
		amount1Str = "5643.17"
		amount1Ts  = time.Now().UTC()
		amount1    = pgtype.Numeric{}
		ts1        = pgtype.Timestamptz{}
	)

	require.NoError(t, amount1.Scan(amount1Str), "failed to parse 5643.17")
	require.NoError(t, ts1.Scan(amount1Ts), "time stamp 1 parse failed.")

	var (
		amount2Str = "-1984.56"
		amount2Ts  = time.Now().Add(time.Minute).UTC()
		amount2    = pgtype.Numeric{}
		ts2        = pgtype.Timestamptz{}
	)

	require.NoError(t, amount2.Scan(amount2Str), "failed to parse -1984.56")
	require.NoError(t, ts2.Scan(amount2Ts), "time stamp 2 parse failed.")

	var (
		amount3Str = "583.81"
		amount3Ts  = time.Now().Add(3 * time.Minute).UTC()
		amount3    = pgtype.Numeric{}
		ts3        = pgtype.Timestamptz{}
	)

	require.NoError(t, amount3.Scan(amount3Str), "failed to parse 583.81")
	require.NoError(t, ts3.Scan(amount3Ts), "time stamp 3 parse failed.")

	// Get general ledger entry test cases.
	testCases := []struct {
		name       string
		expectedTX string
		expectedTS time.Time
		parameter  updateBalanceFiatAccountParams
	}{
		{
			name:       "USD 5643.17",
			expectedTX: amount1Str,
			expectedTS: amount1Ts,
			parameter: updateBalanceFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyUSD,
				LastTx:   amount1,
				LastTxTs: ts1,
			},
		}, {
			name:       "USD -1984.56",
			expectedTX: amount2Str,
			expectedTS: amount2Ts,
			parameter: updateBalanceFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyUSD,
				LastTx:   amount2,
				LastTxTs: ts2,
			},
		}, {
			name:       "USD 583.81",
			expectedTX: amount3Str,
			expectedTS: amount3Ts,
			parameter: updateBalanceFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyUSD,
				LastTx:   amount3,
				LastTxTs: ts3,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)

	defer cancel()

	// Insert new fiat accounts.
	for _, testCase := range testCases {
		test := testCase

		t.Run(fmt.Sprintf("Inserting %s", test.name), func(t *testing.T) {
			result, err := connection.Query.updateBalanceFiatAccount(ctx, &test.parameter)
			require.NoError(t, err, "error expectation condition failed.")

			require.True(t, result.Balance.Valid, "invalid balance.")

			require.True(t, result.LastTx.Valid, "invalid last_tx.")
			driverTX, err := result.LastTx.Value()
			require.NoError(t, err, "failed to get driver value transaction.")
			actualTX, ok := driverTX.(string)
			require.True(t, ok, "failed to extract last_tx string.")
			require.Equal(t, test.expectedTX, actualTX, "expected and actual last_tx mismatched.")

			require.True(t, result.LastTxTs.Valid, "invalid last transaction timestamp.")
			driverTS, err := result.LastTxTs.Value()
			require.NoError(t, err, "failed to get driver value timestamp.")
			actualTS, ok := driverTS.(time.Time)
			require.True(t, ok, "failed to extract last_ts string.")
			require.WithinDuration(t, testCase.expectedTS, actualTS.UTC(), time.Second,
				"expected and actual last_ts mismatched.")
		})
	}

	// Totals check.
	result, err := connection.Query.getFiatAccount(ctx, &getFiatAccountParams{ClientID: clientID1, Currency: CurrencyUSD})
	require.NoError(t, err, "failed to retrieve updated balance.")
	driverValue, err := result.Balance.Value()
	require.NoError(t, err, "failed to get driver value for total.")
	finalBalance, err := strconv.ParseFloat(driverValue.(string), 64)
	require.NoError(t, err, "failed to convert final balance value to float from diver.")

	require.InDelta(t, expectedBalance, finalBalance, 0.01, "expected and actual balance mismatch.")
}

func TestFiat_GeneralLedgerExternalFiatAccount(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, clientID2 := resetTestFiatAccounts(t)

	// Reset the test
	resetTestFiatGeneralLedger(t, clientID1, clientID2)
}

func TestFiat_GeneralLedgerAccountTxFiatAccount(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, clientID2 := resetTestFiatAccounts(t)

	// Reset the test
	resetTestFiatGeneralLedger(t, clientID1, clientID2)

	// Get general ledger entry test cases.
	testCases := []struct {
		name        string
		parameter   generalLedgerAccountTxFiatAccountParams
		errExpected require.ErrorAssertionFunc
	}{
		{
			name: "Client1 - USD",
			parameter: generalLedgerAccountTxFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyUSD,
			},
			errExpected: require.NoError,
		}, {
			name: "Client1 - AED",
			parameter: generalLedgerAccountTxFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyAED,
			},
			errExpected: require.NoError,
		}, {
			name: "Client1 - CAD",
			parameter: generalLedgerAccountTxFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyCAD,
			},
			errExpected: require.NoError,
		}, {
			name: "Client2 - USD",
			parameter: generalLedgerAccountTxFiatAccountParams{
				ClientID: clientID2,
				Currency: CurrencyUSD,
			},
			errExpected: require.NoError,
		}, {
			name: "Client2 - AED",
			parameter: generalLedgerAccountTxFiatAccountParams{
				ClientID: clientID2,
				Currency: CurrencyAED,
			},
			errExpected: require.NoError,
		}, {
			name: "Client2 - CAD",
			parameter: generalLedgerAccountTxFiatAccountParams{
				ClientID: clientID2,
				Currency: CurrencyCAD,
			},
			errExpected: require.NoError,
		}, {
			name: "Client1 - Not Found",
			parameter: generalLedgerAccountTxFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyEUR,
			},
			errExpected: require.NoError,
		},
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)

	defer cancel()

	// Insert new fiat accounts.
	for _, testCase := range testCases {
		test := testCase

		t.Run(fmt.Sprintf("Inserting %s", test.name), func(t *testing.T) {
			results, err := connection.Query.generalLedgerAccountTxFiatAccount(ctx, &test.parameter)
			test.errExpected(t, err, "error expectation condition failed.")
			for _, result := range results {
				require.Equal(t, test.parameter.Currency, result.Currency, "currency type mismatch.")
				require.True(t, result.ClientID.Valid, "invalid UUID.")
				require.True(t, result.TxID.Valid, "invalid TX ID.")
				require.True(t, result.Ammount.Valid, "invalid amount.")
				require.True(t, result.TransactedAt.Valid, "invalid TX time.")
			}
		})
	}
}
