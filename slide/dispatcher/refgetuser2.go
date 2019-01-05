package user

import (
	"context"
)

// START OMIT
type Get struct {
	UserID string
	Result *User
}

func get(ctx context.Context, m *Get) error {
	// ...
}

// END OMIT
