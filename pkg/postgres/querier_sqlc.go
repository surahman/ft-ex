// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Querier interface {
	// purchaseCrypto will execute a transaction to purchase a crypto currency.
	callPurchaseCrypto(ctx context.Context, arg *callPurchaseCryptoParams) error
	// cryptoCreateAccount inserts a fiat account record.
	cryptoCreateAccount(ctx context.Context, arg *cryptoCreateAccountParams) (int64, error)
	// fiatCreateAccount inserts a fiat account record.
	fiatCreateAccount(ctx context.Context, arg *fiatCreateAccountParams) (int64, error)
	// fiatExternalTransferJournalEntry will create both journal entries for fiat accounts inbound deposits.
	fiatExternalTransferJournalEntry(ctx context.Context, arg *fiatExternalTransferJournalEntryParams) (fiatExternalTransferJournalEntryRow, error)
	// fiatGetAccount will retrieve a specific user's account for a given currency.
	fiatGetAccount(ctx context.Context, arg *fiatGetAccountParams) (FiatAccount, error)
	// fiatGetAllAccounts will retrieve all accounts associated with a specific user.
	fiatGetAllAccounts(ctx context.Context, arg *fiatGetAllAccountsParams) ([]FiatAccount, error)
	// fiatGetJournalTransactionForAccountBetweenDates will retrieve the journal entries associated with a specific account
	// in a date range.
	fiatGetAllJournalTransactionPaginated(ctx context.Context, arg *fiatGetAllJournalTransactionPaginatedParams) ([]FiatJournal, error)
	// fiatGetJournalTransaction will retrieve the journal entries associated with a transaction.
	fiatGetJournalTransaction(ctx context.Context, arg *fiatGetJournalTransactionParams) ([]FiatJournal, error)
	// fiatGetJournalTransactionForAccount will retrieve the journal entries associated with a specific account.
	fiatGetJournalTransactionForAccount(ctx context.Context, arg *fiatGetJournalTransactionForAccountParams) ([]FiatJournal, error)
	// fiatInternalTransferJournalEntry will create both journal entries for fiat account internal transfers.
	fiatInternalTransferJournalEntry(ctx context.Context, arg *fiatInternalTransferJournalEntryParams) (fiatInternalTransferJournalEntryRow, error)
	// fiatRowLockAccount will acquire a row level lock without locks on the foreign keys.
	fiatRowLockAccount(ctx context.Context, arg *fiatRowLockAccountParams) (decimal.Decimal, error)
	// fiatUpdateAccountBalance will add an amount to a fiat accounts balance.
	fiatUpdateAccountBalance(ctx context.Context, arg *fiatUpdateAccountBalanceParams) (fiatUpdateAccountBalanceRow, error)
	// testRoundHalfEven
	testRoundHalfEven(ctx context.Context, arg *testRoundHalfEvenParams) (decimal.Decimal, error)
	// userCreate will create a new user record.
	userCreate(ctx context.Context, arg *userCreateParams) (uuid.UUID, error)
	// userDelete will soft delete a users account.
	userDelete(ctx context.Context, clientID uuid.UUID) (int64, error)
	// userGetClientId will retrieve a users client id.
	userGetClientId(ctx context.Context, username string) (uuid.UUID, error)
	// userGetCredentials will retrieve a users client id and password.
	userGetCredentials(ctx context.Context, username string) (userGetCredentialsRow, error)
	// userGetInfo will retrieve a single users account information.
	userGetInfo(ctx context.Context, clientID uuid.UUID) (userGetInfoRow, error)
}

var _ Querier = (*Queries)(nil)
