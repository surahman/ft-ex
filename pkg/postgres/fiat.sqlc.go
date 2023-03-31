// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: fiat.sql

package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const fiatCreateAccount = `-- name: FiatCreateAccount :execrows
INSERT INTO fiat_accounts (client_id, currency)
VALUES ($1, $2)
`

type FiatCreateAccountParams struct {
	ClientID uuid.UUID `json:"clientID"`
	Currency Currency  `json:"currency"`
}

// FiatCreateAccount inserts a fiat account record.
func (q *Queries) FiatCreateAccount(ctx context.Context, arg *FiatCreateAccountParams) (int64, error) {
	result, err := q.db.Exec(ctx, fiatCreateAccount, arg.ClientID, arg.Currency)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const fiatExternalTransferJournalEntry = `-- name: FiatExternalTransferJournalEntry :one
WITH deposit AS (
    INSERT INTO fiat_journal (
        client_id,
        currency,
        amount,
        transacted_at,
        tx_id)
    SELECT
        (   SELECT client_id
            FROM users
            WHERE username = 'deposit-fiat'),
        $2,
        -1 * $3::numeric(18, 2),
        now(),
        gen_random_uuid()
    RETURNING tx_id, transacted_at
)
INSERT INTO fiat_journal (
    client_id,
    currency,
    amount,
    transacted_at,
    tx_id)
SELECT
    $1,
    $2,
    $3::numeric(18, 2),
    (   SELECT transacted_at
        FROM deposit),
    (   SELECT tx_id
        FROM deposit)
RETURNING tx_id, transacted_at
`

type FiatExternalTransferJournalEntryParams struct {
	ClientID uuid.UUID      `json:"clientID"`
	Currency Currency       `json:"currency"`
	Amount   pgtype.Numeric `json:"amount"`
}

type FiatExternalTransferJournalEntryRow struct {
	TxID         uuid.UUID          `json:"txID"`
	TransactedAt pgtype.Timestamptz `json:"transactedAt"`
}

// FiatExternalTransferJournalEntry will create both journal entries for fiat accounts inbound deposits.
func (q *Queries) FiatExternalTransferJournalEntry(ctx context.Context, arg *FiatExternalTransferJournalEntryParams) (FiatExternalTransferJournalEntryRow, error) {
	row := q.db.QueryRow(ctx, fiatExternalTransferJournalEntry, arg.ClientID, arg.Currency, arg.Amount)
	var i FiatExternalTransferJournalEntryRow
	err := row.Scan(&i.TxID, &i.TransactedAt)
	return i, err
}

const fiatGetAccount = `-- name: FiatGetAccount :one
SELECT currency, balance, last_tx, last_tx_ts, created_at, client_id
FROM fiat_accounts
WHERE client_id=$1 AND currency=$2
`

type FiatGetAccountParams struct {
	ClientID uuid.UUID `json:"clientID"`
	Currency Currency  `json:"currency"`
}

// FiatGetAccount will retrieve a specific user's account for a given currency.
func (q *Queries) FiatGetAccount(ctx context.Context, arg *FiatGetAccountParams) (FiatAccount, error) {
	row := q.db.QueryRow(ctx, fiatGetAccount, arg.ClientID, arg.Currency)
	var i FiatAccount
	err := row.Scan(
		&i.Currency,
		&i.Balance,
		&i.LastTx,
		&i.LastTxTs,
		&i.CreatedAt,
		&i.ClientID,
	)
	return i, err
}

const fiatGetAllAccounts = `-- name: FiatGetAllAccounts :many
SELECT currency, balance, last_tx, last_tx_ts, created_at, client_id
FROM fiat_accounts
WHERE client_id=$1
`

// FiatGetAllAccounts will retrieve all accounts associated with a specific user.
func (q *Queries) FiatGetAllAccounts(ctx context.Context, clientID uuid.UUID) ([]FiatAccount, error) {
	rows, err := q.db.Query(ctx, fiatGetAllAccounts, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FiatAccount
	for rows.Next() {
		var i FiatAccount
		if err := rows.Scan(
			&i.Currency,
			&i.Balance,
			&i.LastTx,
			&i.LastTxTs,
			&i.CreatedAt,
			&i.ClientID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fiatGetJournalTransaction = `-- name: FiatGetJournalTransaction :many
SELECT currency, amount, transacted_at, client_id, tx_id
FROM fiat_journal
WHERE tx_id = $1
`

// FiatGetJournalTransaction will retrieve the journal entries associated with a transaction.
func (q *Queries) FiatGetJournalTransaction(ctx context.Context, txID uuid.UUID) ([]FiatJournal, error) {
	rows, err := q.db.Query(ctx, fiatGetJournalTransaction, txID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FiatJournal
	for rows.Next() {
		var i FiatJournal
		if err := rows.Scan(
			&i.Currency,
			&i.Amount,
			&i.TransactedAt,
			&i.ClientID,
			&i.TxID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fiatGetJournalTransactionForAccount = `-- name: FiatGetJournalTransactionForAccount :many
SELECT currency, amount, transacted_at, client_id, tx_id
FROM fiat_journal
WHERE client_id = $1 AND currency = $2
`

type FiatGetJournalTransactionForAccountParams struct {
	ClientID uuid.UUID `json:"clientID"`
	Currency Currency  `json:"currency"`
}

// FiatGetJournalTransactionForAccount will retrieve the journal entries associated with a specific account.
func (q *Queries) FiatGetJournalTransactionForAccount(ctx context.Context, arg *FiatGetJournalTransactionForAccountParams) ([]FiatJournal, error) {
	rows, err := q.db.Query(ctx, fiatGetJournalTransactionForAccount, arg.ClientID, arg.Currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FiatJournal
	for rows.Next() {
		var i FiatJournal
		if err := rows.Scan(
			&i.Currency,
			&i.Amount,
			&i.TransactedAt,
			&i.ClientID,
			&i.TxID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fiatGetJournalTransactionForAccountBetweenDates = `-- name: FiatGetJournalTransactionForAccountBetweenDates :many
SELECT currency, amount, transacted_at, client_id, tx_id
FROM fiat_journal
WHERE client_id = $1
      AND currency = $2
      AND transacted_at
          BETWEEN $3::timestamptz
              AND $4::timestamptz
`

type FiatGetJournalTransactionForAccountBetweenDatesParams struct {
	ClientID  uuid.UUID          `json:"clientID"`
	Currency  Currency           `json:"currency"`
	StartTime pgtype.Timestamptz `json:"startTime"`
	EndTime   pgtype.Timestamptz `json:"endTime"`
}

// FiatGetJournalTransactionForAccountBetweenDates will retrieve the journal entries associated with a specific account
// in a date range.
func (q *Queries) FiatGetJournalTransactionForAccountBetweenDates(ctx context.Context, arg *FiatGetJournalTransactionForAccountBetweenDatesParams) ([]FiatJournal, error) {
	rows, err := q.db.Query(ctx, fiatGetJournalTransactionForAccountBetweenDates,
		arg.ClientID,
		arg.Currency,
		arg.StartTime,
		arg.EndTime,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FiatJournal
	for rows.Next() {
		var i FiatJournal
		if err := rows.Scan(
			&i.Currency,
			&i.Amount,
			&i.TransactedAt,
			&i.ClientID,
			&i.TxID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fiatInternalTransferJournalEntry = `-- name: FiatInternalTransferJournalEntry :one
WITH deposit AS (
    INSERT INTO fiat_journal(
        client_id,
        currency,
        amount,
        transacted_at,
        tx_id)
    SELECT
        $4::uuid,
        $5::currency,
        $6::numeric(18, 2),
        now(),
        gen_random_uuid()
    RETURNING tx_id, transacted_at
)
INSERT INTO fiat_journal (
    client_id,
    currency,
    amount,
    transacted_at,
    tx_id)
SELECT
    $1::uuid,
    $2::currency,
    $3::numeric(18, 2),
    (   SELECT transacted_at
        FROM deposit),
    (   SELECT tx_id
        FROM deposit)
RETURNING tx_id, transacted_at
`

type FiatInternalTransferJournalEntryParams struct {
	DestinationAccount  uuid.UUID      `json:"destinationAccount"`
	DestinationCurrency Currency       `json:"destinationCurrency"`
	CreditAmount        pgtype.Numeric `json:"creditAmount"`
	SourceAccount       uuid.UUID      `json:"sourceAccount"`
	SourceCurrency      Currency       `json:"sourceCurrency"`
	DebitAmount         pgtype.Numeric `json:"debitAmount"`
}

type FiatInternalTransferJournalEntryRow struct {
	TxID         uuid.UUID          `json:"txID"`
	TransactedAt pgtype.Timestamptz `json:"transactedAt"`
}

// FiatInternalTransferJournalEntry will create both journal entries for fiat account internal transfers.
func (q *Queries) FiatInternalTransferJournalEntry(ctx context.Context, arg *FiatInternalTransferJournalEntryParams) (FiatInternalTransferJournalEntryRow, error) {
	row := q.db.QueryRow(ctx, fiatInternalTransferJournalEntry,
		arg.DestinationAccount,
		arg.DestinationCurrency,
		arg.CreditAmount,
		arg.SourceAccount,
		arg.SourceCurrency,
		arg.DebitAmount,
	)
	var i FiatInternalTransferJournalEntryRow
	err := row.Scan(&i.TxID, &i.TransactedAt)
	return i, err
}

const fiatRowLockAccount = `-- name: FiatRowLockAccount :one
SELECT balance
FROM fiat_accounts
WHERE client_id=$1 AND currency=$2
LIMIT 1
FOR NO KEY UPDATE
`

type FiatRowLockAccountParams struct {
	ClientID uuid.UUID `json:"clientID"`
	Currency Currency  `json:"currency"`
}

// FiatRowLockAccount will acquire a row level lock without locks on the foreign keys.
func (q *Queries) FiatRowLockAccount(ctx context.Context, arg *FiatRowLockAccountParams) (pgtype.Numeric, error) {
	row := q.db.QueryRow(ctx, fiatRowLockAccount, arg.ClientID, arg.Currency)
	var balance pgtype.Numeric
	err := row.Scan(&balance)
	return balance, err
}

const fiatUpdateAccountBalance = `-- name: FiatUpdateAccountBalance :one
UPDATE fiat_accounts
SET balance=balance + $3, last_tx=$3, last_tx_ts=$4
WHERE client_id=$1 AND currency=$2
RETURNING balance, last_tx, last_tx_ts
`

type FiatUpdateAccountBalanceParams struct {
	ClientID uuid.UUID          `json:"clientID"`
	Currency Currency           `json:"currency"`
	LastTx   pgtype.Numeric     `json:"lastTx"`
	LastTxTs pgtype.Timestamptz `json:"lastTxTs"`
}

type FiatUpdateAccountBalanceRow struct {
	Balance  pgtype.Numeric     `json:"balance"`
	LastTx   pgtype.Numeric     `json:"lastTx"`
	LastTxTs pgtype.Timestamptz `json:"lastTxTs"`
}

// FiatUpdateAccountBalance will add an amount to a fiat accounts balance.
func (q *Queries) FiatUpdateAccountBalance(ctx context.Context, arg *FiatUpdateAccountBalanceParams) (FiatUpdateAccountBalanceRow, error) {
	row := q.db.QueryRow(ctx, fiatUpdateAccountBalance,
		arg.ClientID,
		arg.Currency,
		arg.LastTx,
		arg.LastTxTs,
	)
	var i FiatUpdateAccountBalanceRow
	err := row.Scan(&i.Balance, &i.LastTx, &i.LastTxTs)
	return i, err
}
