package db

import (
	"context"
	"database/sql"
)

type AutoTxer struct {
	*sql.Tx
	err error
}

func AutoTx(ctx context.Context, db *sql.DB, opts ...*sql.TxOptions) (*AutoTxer, error) {
	if len(opts) != 0 {
		tx, err := db.BeginTx(ctx, opts[0])
		if err != nil {
			return nil, err
		}
		return &AutoTxer{Tx: tx}, nil
	}
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return &AutoTxer{Tx: tx}, nil
}

func (at *AutoTxer) Close() error {
	if at.err != nil {
		return at.Rollback()
	}
	return at.Commit()
}

func (at *AutoTxer) SetErr(err error) {
	at.err = err
}

func (at *AutoTxer) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := at.Tx.ExecContext(ctx, query, args...)
	at.err = err
	return result, err
}

func (at *AutoTxer) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	result, err := at.Tx.QueryContext(ctx, query, args...)
	at.err = err
	return result, err
}
