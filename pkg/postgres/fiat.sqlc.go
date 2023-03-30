// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: fiat.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const fiatCreateAccount = `-- name: fiatCreateAccount :execrows
INSERT INTO fiat_accounts (client_id, currency)
VALUES ($1, $2)
`

type fiatCreateAccountParams struct {
	ClientID pgtype.UUID `json:"clientID"`
	Currency Currency    `json:"currency"`
}

// fiatCreateAccount inserts a fiat account record.
func (q *Queries) fiatCreateAccount(ctx context.Context, arg *fiatCreateAccountParams) (int64, error) {
	result, err := q.db.Exec(ctx, fiatCreateAccount, arg.ClientID, arg.Currency)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const fiatExternalTransferJournalEntry = `-- name: fiatExternalTransferJournalEntry :one
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

type fiatExternalTransferJournalEntryParams struct {
	ClientID pgtype.UUID    `json:"clientID"`
	Currency Currency       `json:"currency"`
	Amount   pgtype.Numeric `json:"amount"`
}

type fiatExternalTransferJournalEntryRow struct {
	TxID         pgtype.UUID        `json:"txID"`
	TransactedAt pgtype.Timestamptz `json:"transactedAt"`
}

// fiatExternalTransferJournalEntry will create both journal entries for fiat accounts inbound deposits.
func (q *Queries) fiatExternalTransferJournalEntry(ctx context.Context, arg *fiatExternalTransferJournalEntryParams) (fiatExternalTransferJournalEntryRow, error) {
	row := q.db.QueryRow(ctx, fiatExternalTransferJournalEntry, arg.ClientID, arg.Currency, arg.Amount)
	var i fiatExternalTransferJournalEntryRow
	err := row.Scan(&i.TxID, &i.TransactedAt)
	return i, err
}

const fiatGetAccount = `-- name: fiatGetAccount :one
SELECT currency, balance, last_tx, last_tx_ts, created_at, client_id
FROM fiat_accounts
WHERE client_id=$1 AND currency=$2
`

type fiatGetAccountParams struct {
	ClientID pgtype.UUID `json:"clientID"`
	Currency Currency    `json:"currency"`
}

// getFiatAccount will retrieve a specific user's account for a given currency.
func (q *Queries) fiatGetAccount(ctx context.Context, arg *fiatGetAccountParams) (FiatAccount, error) {
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

const fiatGetAllAccounts = `-- name: fiatGetAllAccounts :many
SELECT currency, balance, last_tx, last_tx_ts, created_at, client_id
FROM fiat_accounts
WHERE client_id=$1
`

// fiatGetAllAccounts will retrieve all accounts associated with a specific user.
func (q *Queries) fiatGetAllAccounts(ctx context.Context, clientID pgtype.UUID) ([]FiatAccount, error) {
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

const fiatGetJournalTransaction = `-- name: fiatGetJournalTransaction :many
SELECT currency, amount, transacted_at, client_id, tx_id
FROM fiat_journal
WHERE tx_id = $1
`

// fiatGetJournalTransaction will retrieve the journal entries associated with a transaction.
func (q *Queries) fiatGetJournalTransaction(ctx context.Context, txID pgtype.UUID) ([]FiatJournal, error) {
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

const fiatGetJournalTransactionForAccount = `-- name: fiatGetJournalTransactionForAccount :many
SELECT currency, amount, transacted_at, client_id, tx_id
FROM fiat_journal
WHERE client_id = $1 AND currency = $2
`

type fiatGetJournalTransactionForAccountParams struct {
	ClientID pgtype.UUID `json:"clientID"`
	Currency Currency    `json:"currency"`
}

// fiatGetJournalTransactionForAccount will retrieve the journal entries associated with a specific account.
func (q *Queries) fiatGetJournalTransactionForAccount(ctx context.Context, arg *fiatGetJournalTransactionForAccountParams) ([]FiatJournal, error) {
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

const fiatGetJournalTransactionForAccountBetweenDates = `-- name: fiatGetJournalTransactionForAccountBetweenDates :many
SELECT currency, amount, transacted_at, client_id, tx_id
FROM fiat_journal
WHERE client_id = $1
      AND currency = $2
      AND transacted_at
          BETWEEN $3::timestamptz
              AND $4::timestamptz
`

type fiatGetJournalTransactionForAccountBetweenDatesParams struct {
	ClientID  pgtype.UUID        `json:"clientID"`
	Currency  Currency           `json:"currency"`
	StartTime pgtype.Timestamptz `json:"startTime"`
	EndTime   pgtype.Timestamptz `json:"endTime"`
}

// fiatGetJournalTransactionForAccountBetweenDates will retrieve the journal entries associated with a specific account
// in a date range.
func (q *Queries) fiatGetJournalTransactionForAccountBetweenDates(ctx context.Context, arg *fiatGetJournalTransactionForAccountBetweenDatesParams) ([]FiatJournal, error) {
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

const fiatInternalTransferJournalEntry = `-- name: fiatInternalTransferJournalEntry :one
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

type fiatInternalTransferJournalEntryParams struct {
	DestinationAccount  pgtype.UUID    `json:"destinationAccount"`
	DestinationCurrency Currency       `json:"destinationCurrency"`
	CreditAmount        pgtype.Numeric `json:"creditAmount"`
	SourceAccount       pgtype.UUID    `json:"sourceAccount"`
	SourceCurrency      Currency       `json:"sourceCurrency"`
	DebitAmount         pgtype.Numeric `json:"debitAmount"`
}

type fiatInternalTransferJournalEntryRow struct {
	TxID         pgtype.UUID        `json:"txID"`
	TransactedAt pgtype.Timestamptz `json:"transactedAt"`
}

// fiatInternalTransferJournalEntry will create both journal entries for fiat account internal transfers.
func (q *Queries) fiatInternalTransferJournalEntry(ctx context.Context, arg *fiatInternalTransferJournalEntryParams) (fiatInternalTransferJournalEntryRow, error) {
	row := q.db.QueryRow(ctx, fiatInternalTransferJournalEntry,
		arg.DestinationAccount,
		arg.DestinationCurrency,
		arg.CreditAmount,
		arg.SourceAccount,
		arg.SourceCurrency,
		arg.DebitAmount,
	)
	var i fiatInternalTransferJournalEntryRow
	err := row.Scan(&i.TxID, &i.TransactedAt)
	return i, err
}

const fiatRowLockAccount = `-- name: fiatRowLockAccount :one
SELECT balance
FROM fiat_accounts
WHERE client_id=$1 AND currency=$2
LIMIT 1
FOR NO KEY UPDATE
`

type fiatRowLockAccountParams struct {
	ClientID pgtype.UUID `json:"clientID"`
	Currency Currency    `json:"currency"`
}

// fiatRowLockAccount will acquire a row level lock without locks on the foreign keys.
func (q *Queries) fiatRowLockAccount(ctx context.Context, arg *fiatRowLockAccountParams) (pgtype.Numeric, error) {
	row := q.db.QueryRow(ctx, fiatRowLockAccount, arg.ClientID, arg.Currency)
	var balance pgtype.Numeric
	err := row.Scan(&balance)
	return balance, err
}

const fiatUpdateAccountBalance = `-- name: fiatUpdateAccountBalance :one
UPDATE fiat_accounts
SET balance=balance + $3, last_tx=$3, last_tx_ts=$4
WHERE client_id=$1 AND currency=$2
RETURNING balance, last_tx, last_tx_ts
`

type fiatUpdateAccountBalanceParams struct {
	ClientID pgtype.UUID        `json:"clientID"`
	Currency Currency           `json:"currency"`
	LastTx   pgtype.Numeric     `json:"lastTx"`
	LastTxTs pgtype.Timestamptz `json:"lastTxTs"`
}

type fiatUpdateAccountBalanceRow struct {
	Balance  pgtype.Numeric     `json:"balance"`
	LastTx   pgtype.Numeric     `json:"lastTx"`
	LastTxTs pgtype.Timestamptz `json:"lastTxTs"`
}

// fiatUpdateAccountBalance will add an amount to a fiat accounts balance.
func (q *Queries) fiatUpdateAccountBalance(ctx context.Context, arg *fiatUpdateAccountBalanceParams) (fiatUpdateAccountBalanceRow, error) {
	row := q.db.QueryRow(ctx, fiatUpdateAccountBalance,
		arg.ClientID,
		arg.Currency,
		arg.LastTx,
		arg.LastTxTs,
	)
	var i fiatUpdateAccountBalanceRow
	err := row.Scan(&i.Balance, &i.LastTx, &i.LastTxTs)
	return i, err
}
