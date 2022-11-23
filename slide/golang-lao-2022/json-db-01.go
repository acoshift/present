package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// START 1 OMIT
type Detail struct {
	//
}

func (d Detail) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *Detail) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var b []byte
	switch p := src.(type) {
	case []byte:
		b = p
	case string:
		b = []byte(p)
	default:
		return fmt.Errorf("src type not support")
	}
	return json.Unmarshal(b, d)
}

// END 1 OMIT

// START 2 OMIT
func getDetail(ctx context.Context, db *sql.DB, id int64) (*Detail, error) {
	var d Detail
	err := db.QueryRowContext(ctx, `
		select detail
		from table
		where id = $1
	`, id).Scan(&d)
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
	`, id, detail)
	return err
}

// END 3 OMIT
