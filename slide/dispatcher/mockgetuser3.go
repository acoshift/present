package user

import "context"

type UserStorage interface {
	Create(ctx context.Context, user *User) (string, error)
	Get(ctx context.Context, userID string) (*User, error)
	SetName(ctx context.Context, userID string, name string) error
	SetActive(ctx context.Context, userID string, active bool) error
	Delete(ctx context.Context, userID string) error
}

// START OMIT
type MockUserStorage struct {
	CreateFunc    func(ctx context.Context, user *User) (string, error)
	GetFunc       func(ctx context.Context, userID string) (*User, error)
	SetNameFunc   func(ctx context.Context, userID string, name string) error
	SetActiveFunc func(ctx context.Context, userID string, active bool) error
	DeleteFunc    func(ctx context.Context, userID string) error
}

func (m *MockUserStorage) Create(ctx context.Context, user *User) (string, error) {
	return m.CreateFunc(ctx, user)
}

func (m *MockUserStorage) Get(ctx context.Context, userID string) (*User, error) {
	return m.GetFunc(ctx, userID)
}

func (m *MockUserStorage) SetName(ctx context.Context, userID string, name string) error {
	return m.SetNameFunc(ctx, userID, name)
}

func (m *MockUserStorage) SetActive(ctx context.Context, userID string, active bool) error {
	return m.SetActiveFunc(ctx, userID, active)
}

func (m *MockUserStorage) Delete(ctx context.Context, userID string) error {
	return m.DeleteFunc(ctx, userID)
}

// END OMIT
