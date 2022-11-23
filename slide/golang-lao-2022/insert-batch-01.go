package main

import (
	"context"
	"database/sql"

	"github.com/acoshift/pgsql/pgstmt"
)

var db *sql.DB

type Tx struct {
	UserID string
	Amount int64
	Type   int
}

// START 1 OMIT
func insertTx(ctx context.Context, tx *Tx) error {
	_, err := db.ExecContext(ctx, `
		insert into txs (user_id, amount, type)
		values ($1, $2, $3)
	`, tx.UserID, tx.Amount, tx.Type)
	return err
}

// END 1 OMIT

// START 2 OMIT
func insertTxs(ctx context.Context, txs []*Tx) error {
	_, err := pgstmt.Insert(func(b pgstmt.InsertStatement) {
		b.Into("txs")
		b.Columns("user_id", "amount", "type")
		for _, tx := range txs {
			b.Values(tx.UserID, tx.Amount, tx.Type)
		}
	}).ExecContext(ctx, db.ExecContext)
	return err
}

// END 2 OMIT
