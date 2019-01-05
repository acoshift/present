package user

import (
	"context"
	"database/sql"
)

type User struct {
	ID   string
	Name string
}

func Get(ctx context.Context, userID string) (*User, error) {
	var user User

	q := getQueryer(ctx)
	err := q.QueryRowContext(ctx, `
		select
			id, name
		from users
		where id = $1
	`, userID).Scan(
		&user.ID, &user.Name,
	)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
