package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

type Detail struct {
	//
}

// START 1 OMIT
type jsonWrapper struct {
	value any
}

func (w jsonWrapper) Value() (driver.Value, error) {
	//
}

func (w *jsonWrapper) Scan(src interface{}) error {
	//
}

func wrapJSON(v any) interface {
	driver.Valuer
	sql.Scanner
} {
	return &jsonWrapper{v}
}

// END 1 OMIT

// START 2 OMIT
func getDetail(ctx context.Context, db *sql.DB, id int64) (*Detail, error) {
	var d Detail
	err := db.QueryRowContext(ctx, `
		select detail
		from table
		where id = $1
	`, id).Scan(wrapJSON(&d))
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// END 2 OMIT

// START 3 OMIT
func saveDetail(ctx context.Context, db *sql.DB, id int64, detail *Detail) error {
	_, err := db.ExecContext(ctx, `
		update table
		set detail = $2
		where id = $1
	`, id, wrapJSON(&detail))
	return err
}

// END 3 OMIT
