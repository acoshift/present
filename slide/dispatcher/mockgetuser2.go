package user

import "context"

// START OMIT
type UserStorage interface {
	Create(ctx context.Context, user *User) (string, error)
	Get(ctx context.Context, userID string) (*User, error)
	SetName(ctx context.Context, userID string, name string) error
	SetActive(ctx context.Context, userID string, active bool) error
	Delete(ctx context.Context, userID string) error
}

// END OMIT
