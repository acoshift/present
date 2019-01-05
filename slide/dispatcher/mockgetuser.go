package user

import "context"

// START OMIT
type UserStorage interface {
	Get(ctx context.Context, userID string) (*User, error)
}

type storage struct{}

func (storage) Get(ctx context.Context, userID string) (*User, error) {
	// ...
}

// END OMIT
