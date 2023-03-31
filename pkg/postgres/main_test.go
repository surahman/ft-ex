package postgres

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"github.com/surahman/FTeX/pkg/constants"
	"github.com/surahman/FTeX/pkg/logger"
	"go.uber.org/zap"
)

// postgresConfigTestData is a map Postgres configuration test data.
var postgresConfigTestData = configTestData()

// configFileKey is the name of the Postgres configuration file to use in the tests.
var configFileKey string

// connection pool to Cassandra cluster.
var connection *Postgres

// zapLogger is the Zap logger used strictly for the test suite in this package.
var zapLogger *logger.Logger

func TestMain(m *testing.M) {
	// Parse commandline flags to check for short tests.
	flag.Parse()

	var err error
	// Configure logger.
	if zapLogger, err = logger.NewTestLogger(); err != nil {
		log.Printf("Test suite logger setup failed: %v\n", err)
		os.Exit(1)
	}

	// Setup test space.
	if err = setup(); err != nil {
		zapLogger.Error("Test suite setup failure", zap.Error(err))
		os.Exit(1)
	}

	// Run test suite.
	exitCode := m.Run()

	// Cleanup test space.
	if err = tearDown(); err != nil {
		zapLogger.Error("Test suite teardown failure:", zap.Error(err))
		os.Exit(1)
	}

	os.Exit(exitCode)
}

// setup will configure the connection to the test database.
func setup() error {
	if testing.Short() {
		zapLogger.Warn("Short test: Skipping Postgres integration tests")

		return nil
	}

	var err error

	// If running on a GitHub Actions runner use the default credentials for Postgres.
	configFileKey = "test_suite"
	if _, ok := os.LookupEnv(constants.GetGithubCIKey()); ok == true {
		configFileKey = "github-ci-runner"

		zapLogger.Info("Integration Test running on Github CI runner.")
	}

	// Setup mock filesystem for configs.
	fs := afero.NewMemMapFs()
	if err = fs.MkdirAll(constants.GetEtcDir(), 0644); err != nil {
		return fmt.Errorf("afero memory mapped file system setup failed: %w", err)
	}

	if err = afero.WriteFile(fs, constants.GetEtcDir()+constants.GetPostgresFileName(),
		[]byte(postgresConfigTestData[configFileKey]), 0644); err != nil {
		return fmt.Errorf("afero memory mapped file system write failed: %w", err)
	}

	// Load Postgres configurations for test suite.
	if connection, err = NewPostgres(&fs, zapLogger); err != nil {
		return err
	}

	if err = connection.Open(); err != nil {
		return fmt.Errorf("postgres connection opening failed: %w", err)
	}

	return nil
}

// tearDown will delete the test clusters keyspace.
func tearDown() (err error) {
	if !testing.Short() {
		if err := connection.Close(); err != nil {
			return fmt.Errorf("postgres connection termination failure in test suite: %w", err)
		}
	}

	return
}

// insertTestUsers will reset the users table and create some test user accounts.
func insertTestUsers(t *testing.T) {
	t.Helper()

	// Reset the users table.
	query := "DELETE FROM users WHERE first_name != 'Internal';"
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)

	defer cancel()

	rows, err := connection.Query.db.Query(ctx, query)
	rows.Close()

	require.NoError(t, err, "failed to wipe users table before reinserting users.")

	// Insert new users.
	for key, testCase := range getTestUsers() {
		user := testCase

		t.Run(fmt.Sprintf("Inserting %s", key), func(t *testing.T) {
			clientID, err := connection.Query.UserCreate(ctx, &user)
			require.NoErrorf(t, err, "failed to insert test user account: %w", err)
			require.True(t, clientID.Valid, "failed to retrieve client id from response")
		})
	}
}

// resetTestFiatAccounts will reset the fiat accounts table and create some test accounts.
func resetTestFiatAccounts(t *testing.T) (pgtype.UUID, pgtype.UUID) {
	t.Helper()

	// Reset the fiat accounts table.
	query := "TRUNCATE TABLE fiat_accounts CASCADE;"
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)

	defer cancel()

	rows, err := connection.Query.db.Query(ctx, query)
	rows.Close()

	require.NoError(t, err, "failed to wipe fiat accounts table before reinserting accounts.")

	// Retrieve client ids from users table.
	clientID1, err := connection.Query.UserGetClientId(ctx, "username1")
	require.NoError(t, err, "failed to retrieve username1 client id.")
	clientID2, err := connection.Query.UserGetClientId(ctx, "username2")
	require.NoError(t, err, "failed to retrieve username2 client id.")

	// Insert new fiat accounts.
	for key, testCase := range getTestFiatAccounts(clientID1, clientID2) {
		parameters := testCase

		t.Run(fmt.Sprintf("Inserting %s", key), func(t *testing.T) {
			for _, param := range parameters {
				accInfo := param
				rowCount, err := connection.Query.FiatCreateAccount(ctx, &accInfo)
				require.NoError(t, err, "errored whilst trying to insert fiat account.")
				require.NotEqual(t, 0, rowCount, "no rows were added.")
			}
		})
	}

	return clientID1, clientID2
}

// resetTestFiatJournal will reset the fiat journal with base internal and external entries.
func resetTestFiatJournal(t *testing.T, clientID1, clientID2 pgtype.UUID) {
	t.Helper()

	// Reset the fiat journal table.
	query := "TRUNCATE TABLE fiat_journal CASCADE;"
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)

	defer cancel()

	rows, err := connection.Query.db.Query(ctx, query)
	rows.Close()

	require.NoError(t, err, "failed to wipe fiat journal table before reinserting entries.")

	// Get general ledger entry test cases.
	testCases, err := getTestFiatJournal(clientID1, clientID2)
	require.NoError(t, err, "failed to generate test cases.")

	// Insert new fiat accounts.
	for key, testCase := range testCases {
		parameters := testCase

		t.Run(fmt.Sprintf("Inserting %s", key), func(t *testing.T) {
			result, err := connection.Query.FiatExternalTransferJournalEntry(ctx, &parameters)
			require.NoError(t, err, "failed to insert external fiat account entry.")
			require.True(t, result.TxID.Valid, "returned transaction id is invalid.")
			require.True(t, result.TransactedAt.Valid, "returned transaction time is invalid.")
		})
	}
}

// insertTestInternalFiatGeneralLedger will not reset the journal and will insert some test internal transfers.
func insertTestInternalFiatGeneralLedger(t *testing.T, clientID1, clientID2 pgtype.UUID) (
	map[string]FiatInternalTransferJournalEntryParams, map[string]FiatInternalTransferJournalEntryRow) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)

	defer cancel()

	// Get journal entry test cases.
	testCases, err := getTestJournalInternalFiatAccounts(clientID1, clientID2)
	require.NoError(t, err, "failed to generate test cases.")

	// Mapping for transactions to parameters.
	transactions := make(map[string]FiatInternalTransferJournalEntryRow, len(testCases))

	// Insert new fiat accounts.
	for key, testCase := range testCases {
		parameters := testCase

		t.Run(fmt.Sprintf("Inserting %s", key), func(t *testing.T) {
			row, err := connection.Query.FiatInternalTransferJournalEntry(ctx, &parameters)
			require.NoError(t, err, "errored whilst inserting internal fiat general ledger entry.")
			require.NotEqual(t, 0, row, "no rows were added.")
			transactions[key] = row
		})
	}

	return testCases, transactions
}
