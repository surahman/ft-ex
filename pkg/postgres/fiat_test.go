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
		amount1Ts = time.Now().UTC()
		amount1   = pgtype.Numeric{}
		ts1       = pgtype.Timestamptz{}
	)

	require.NoError(t, amount1.Scan("5643.17"), "failed to parse 5643.17")
	require.NoError(t, ts1.Scan(amount1Ts), "time stamp 1 parse failed.")

	var (
		amount2Ts = time.Now().Add(time.Minute).UTC()
		amount2   = pgtype.Numeric{}
		ts2       = pgtype.Timestamptz{}
	)

	require.NoError(t, amount2.Scan("-1984.56"), "failed to parse -1984.56")
	require.NoError(t, ts2.Scan(amount2Ts), "time stamp 2 parse failed.")

	var (
		amount3Ts = time.Now().Add(3 * time.Minute).UTC()
		amount3   = pgtype.Numeric{}
		ts3       = pgtype.Timestamptz{}
	)

	require.NoError(t, amount3.Scan("583.81"), "failed to parse 583.81")
	require.NoError(t, ts3.Scan(amount3Ts), "time stamp 3 parse failed.")

	// Get general ledger entry test cases.
	testCases := []struct {
		name       string
		expectedTS time.Time
		parameter  updateBalanceFiatAccountParams
	}{
		{
			name:       "USD 5643.17",
			expectedTS: amount1Ts,
			parameter: updateBalanceFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyUSD,
				LastTx:   amount1,
				LastTxTs: ts1,
			},
		}, {
			name:       "USD -1984.56",
			expectedTS: amount2Ts,
			parameter: updateBalanceFiatAccountParams{
				ClientID: clientID1,
				Currency: CurrencyUSD,
				LastTx:   amount2,
				LastTxTs: ts2,
			},
		}, {
			name:       "USD 583.81",
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
			require.Equal(t, test.parameter.LastTx, result.LastTx, "expected and actual last_tx mismatched.")

			require.True(t, result.LastTxTs.Valid, "invalid last transaction timestamp.")
			require.WithinDuration(t, test.expectedTS, result.LastTxTs.Time, time.Second,
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

func TestFiat_FiatExternalTransferJournalEntry(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, clientID2 := resetTestFiatAccounts(t)

	// Reset the External Fiat General Ledger.
	resetTestFiatJournal(t, clientID1, clientID2)
}

func TestFiat_FiatInternalTransferJournalEntry(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, clientID2 := resetTestFiatAccounts(t)

	// Reset the External Fiat Journal.
	resetTestFiatJournal(t, clientID1, clientID2)

	// Insert internal fiat journal transactions.
	_, _ = insertTestInternalFiatGeneralLedger(t, clientID1, clientID2)
}

func TestFiat_FiatGetJournalTransaction(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, clientID2 := resetTestFiatAccounts(t)

	// Reset the external fiat journal entries.
	resetTestFiatJournal(t, clientID1, clientID2)

	// Insert internal fiat journal transactions.
	testCases, txRows := insertTestInternalFiatGeneralLedger(t, clientID1, clientID2)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)

	defer cancel()

	for key, row := range txRows {
		param := row

		t.Run(fmt.Sprintf("Retrieving %s", key), func(t *testing.T) {
			tx := testCases[key]

			result, err := connection.Query.fiatGetJournalTransaction(ctx, param.TxID)
			require.NoError(t, err, "error expectation condition failed.")
			require.Equal(t, 2, len(result), "incorrect row count returned.")

			var (
				srcRecord = result[0]
				dstRecord = result[1]
			)

			if srcRecord.Currency != tx.SourceCurrency {
				srcRecord = result[1]
				dstRecord = result[0]
			}

			require.Equal(t, srcRecord.Currency, tx.SourceCurrency, "source currency mismatch.")
			require.Equal(t, dstRecord.Currency, tx.DestinationCurrency, "destination currency mismatch.")
			require.Equal(t, srcRecord.ClientID, tx.SourceAccount, "source client id mismatch.")
			require.Equal(t, dstRecord.ClientID, tx.DestinationAccount, "destination client id mismatch.")
			require.Equal(t, srcRecord.Ammount, tx.DebitAmount, "source amount mismatch.")
			require.Equal(t, dstRecord.Ammount, tx.CreditAmount, "destination amount mismatch.")
		})
	}
}

func TestFiat_FiatGetJournalTransactionForAccount(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, clientID2 := resetTestFiatAccounts(t)

	// Reset the test
	resetTestFiatJournal(t, clientID1, clientID2)

	// Get journal entry test cases.
	testCases := []struct {
		name        string
		parameter   fiatGetJournalTransactionForAccountParams
		errExpected require.ErrorAssertionFunc
	}{
		{
			name: "Client1 - USD",
			parameter: fiatGetJournalTransactionForAccountParams{
				ClientID: clientID1,
				Currency: CurrencyUSD,
			},
			errExpected: require.NoError,
		}, {
			name: "Client1 - AED",
			parameter: fiatGetJournalTransactionForAccountParams{
				ClientID: clientID1,
				Currency: CurrencyAED,
			},
			errExpected: require.NoError,
		}, {
			name: "Client1 - CAD",
			parameter: fiatGetJournalTransactionForAccountParams{
				ClientID: clientID1,
				Currency: CurrencyCAD,
			},
			errExpected: require.NoError,
		}, {
			name: "Client2 - USD",
			parameter: fiatGetJournalTransactionForAccountParams{
				ClientID: clientID2,
				Currency: CurrencyUSD,
			},
			errExpected: require.NoError,
		}, {
			name: "Client2 - AED",
			parameter: fiatGetJournalTransactionForAccountParams{
				ClientID: clientID2,
				Currency: CurrencyAED,
			},
			errExpected: require.NoError,
		}, {
			name: "Client2 - CAD",
			parameter: fiatGetJournalTransactionForAccountParams{
				ClientID: clientID2,
				Currency: CurrencyCAD,
			},
			errExpected: require.NoError,
		}, {
			name: "Client1 - Not Found",
			parameter: fiatGetJournalTransactionForAccountParams{
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
			results, err := connection.Query.fiatGetJournalTransactionForAccount(ctx, &test.parameter)
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

func TestFiat_GetFiatAccount(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, clientID2 := resetTestFiatAccounts(t)

	// Reset the test
	resetTestFiatJournal(t, clientID1, clientID2)

	// Test grid.
	testCases := []struct {
		name            string
		parameter       getFiatAccountParams
		errExpectation  require.ErrorAssertionFunc
		boolExpectation require.BoolAssertionFunc
	}{
		{
			name: "ClientID 1 - Not found",
			parameter: getFiatAccountParams{
				ClientID: clientID1,
				Currency: "PKR",
			},
			errExpectation:  require.Error,
			boolExpectation: require.False,
		}, {
			name: "ClientID 1 - USD",
			parameter: getFiatAccountParams{
				ClientID: clientID1,
				Currency: "USD",
			},
			errExpectation:  require.NoError,
			boolExpectation: require.True,
		}, {
			name: "ClientID 1 - CAD",
			parameter: getFiatAccountParams{
				ClientID: clientID1,
				Currency: "CAD",
			},
			errExpectation:  require.NoError,
			boolExpectation: require.True,
		}, {
			name: "ClientID 1 - AED",
			parameter: getFiatAccountParams{
				ClientID: clientID1,
				Currency: "AED",
			},
			errExpectation:  require.NoError,
			boolExpectation: require.True,
		},
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)

	defer cancel()

	// Insert new fiat accounts.
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Retrieving %s", testCase.name), func(t *testing.T) {
			results, err := connection.Query.getFiatAccount(ctx, &testCase.parameter)
			testCase.errExpectation(t, err, "error expectation failed.")
			testCase.boolExpectation(t, results.ClientID.Valid, "clientId validity expectation failed.")
			testCase.boolExpectation(t, results.LastTxTs.Valid, "lastTxTs validity expectation failed.")
			testCase.boolExpectation(t, results.Balance.Valid, "balance validity expectation failed.")
			testCase.boolExpectation(t, results.LastTx.Valid, "lastTx validity expectation failed.")
			testCase.boolExpectation(t, results.CreatedAt.Valid, "createdAt validity expectation failed.")
			testCase.boolExpectation(t, results.Currency.Valid(), "currency validity expectation failed.")

			if err != nil {
				return
			}

			require.Equal(t, testCase.parameter.Currency, results.Currency, "currency mismatch.")
			require.Equal(t, testCase.parameter.ClientID, results.ClientID, "clientId mismatch.")
		})
	}
}

func TestFiat_GetAllFiatAccounts(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, _ := resetTestFiatAccounts(t)

	// Testing grid.
	testCases := []struct {
		name           string
		clientID       pgtype.UUID
		expectedRowCnt int
	}{
		{
			name:           "ClientID 1",
			clientID:       clientID1,
			expectedRowCnt: 3,
		}, {
			name:           "Nonexistent",
			clientID:       pgtype.UUID{},
			expectedRowCnt: 0,
		},
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)

	defer cancel()

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Retrieving %s", testCase.name), func(t *testing.T) {
			rows, err := connection.Query.getAllFiatAccounts(ctx, testCase.clientID)
			require.NoError(t, err, "error expectation failed.")
			require.Equal(t, testCase.expectedRowCnt, len(rows), "expected row count mismatch.")
		})
	}
}

func TestFiat_FiatGetJournalTransactionForAccountBetweenDates(t *testing.T) {
	// Skip integration tests for short test runs.
	if testing.Short() {
		return
	}

	// Insert test users.
	insertTestUsers(t)

	// Insert initial set of test fiat accounts.
	clientID1, clientID2 := resetTestFiatAccounts(t)

	// Reset the test
	resetTestFiatJournal(t, clientID1, clientID2)

	// Context setup for no hold-and-wait.
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)

	defer cancel()

	// Insert some more fiat journal entries for good measure.
	{
		parameters, err := getTestFiatJournal(clientID1, clientID2)
		require.NoError(t, err, "failed to get parameters to insert additional fiat journal entries.")
		for _, item := range parameters {
			parameter := item
			t.Run(fmt.Sprintf("Inserting %v - %s", parameter.ClientID, parameter.Currency), func(t *testing.T) {
				for idx := 0; idx < 3; idx++ {
					_, err := connection.Query.fiatExternalTransferJournalEntry(ctx, &parameter)
					require.NoError(t, err, "error expectation failed.")
				}
			})
		}
	}

	// Setup time intervals.
	var (
		timePoint    = time.Now().UTC()
		minuteAhead  = pgtype.Timestamptz{}
		minuteBehind = pgtype.Timestamptz{}
		hourAhead    = pgtype.Timestamptz{}
		hourBehind   = pgtype.Timestamptz{}
	)

	require.NoError(t, minuteAhead.Scan(timePoint.Add(time.Minute)))
	require.NoError(t, minuteBehind.Scan(timePoint.Add(-time.Minute)))
	require.NoError(t, hourAhead.Scan(timePoint.Add(time.Hour)))
	require.NoError(t, hourBehind.Scan(timePoint.Add(-time.Hour)))

	// Test grid.
	testCases := []struct {
		name         string
		expectedCont int
		parameters   fiatGetJournalTransactionForAccountBetweenDatesParams
	}{
		{
			name:         "ClientID1 USD: Before-After",
			expectedCont: 4,
			parameters: fiatGetJournalTransactionForAccountBetweenDatesParams{
				ClientID:  clientID1,
				Currency:  "USD",
				StartTime: minuteBehind,
				EndTime:   minuteAhead,
			},
		}, {
			name:         "ClientID1 USD: Before",
			expectedCont: 0,
			parameters: fiatGetJournalTransactionForAccountBetweenDatesParams{
				ClientID:  clientID1,
				Currency:  "USD",
				StartTime: hourBehind,
				EndTime:   minuteBehind,
			},
		}, {
			name:         "ClientID1 USD: After",
			expectedCont: 0,
			parameters: fiatGetJournalTransactionForAccountBetweenDatesParams{
				ClientID:  clientID1,
				Currency:  "USD",
				StartTime: minuteAhead,
				EndTime:   hourAhead,
			},
		}, {
			name:         "ClientID2 - AED: Before-After",
			expectedCont: 4,
			parameters: fiatGetJournalTransactionForAccountBetweenDatesParams{
				ClientID:  clientID2,
				Currency:  "AED",
				StartTime: minuteBehind,
				EndTime:   minuteAhead,
			},
		}, {
			name:         "ClientID2 - PKR: Before-After",
			expectedCont: 0,
			parameters: fiatGetJournalTransactionForAccountBetweenDatesParams{
				ClientID:  clientID2,
				Currency:  "PKR",
				StartTime: minuteBehind,
				EndTime:   minuteAhead,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Retrieving %s", testCase.name), func(t *testing.T) {
			rows, err := connection.Query.fiatGetJournalTransactionForAccountBetweenDates(ctx, &testCase.parameters)
			require.NoError(t, err, "error expectation failed.")
			require.Equal(t, testCase.expectedCont, len(rows), "expected row count mismatch.")
		})
	}
}
