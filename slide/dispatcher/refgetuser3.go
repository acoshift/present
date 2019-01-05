package user

// START OMIT
type Create struct {
	User   *User
	Result string // User ID
}

type Get struct {
	UserID string
	Result *User
}

type SetName struct {
	UserID string
	Name   string
}

type Delete struct {
	UserID string
}

// END OMIT
