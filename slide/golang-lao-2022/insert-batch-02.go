package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/acoshift/pgsql/pgstmt"
)

type Tx struct{}

// START 1 OMIT
var insertCh = make(chan *Tx, 1000)

func insertWorker(ctx context.Context) {
	const batchSize = 1000
	buf := make([]*Tx, 0, 1000)
	flush := func() {
		if len(buf) == 0 {
			return
		}
		err := insertTxs(ctx, buf)
		if err != nil {
			log.Printf("batch insert error; %v", err)
			return
		}
		buf = buf[:0]
	}
	// END 1 OMIT
	// START 2 OMIT

	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			flush()
			return
		case tx := <-insertCh:
			buf = append(buf, tx)
			if len(buf) >= batchSize {
				flush()
			}
		case <-ticker.C:
			flush()
		}
	}
}

// END 2 OMIT

func insertTx(tx *Tx) {
	insertCh <- tx
}

var db *sql.DB

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
