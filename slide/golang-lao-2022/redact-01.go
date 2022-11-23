package main

import (
	"encoding/json"
	"os"
)

// START OMIT
func main() {
	req := SignInParams{
		Username: "admin",
		Password: "123456",
	}
	json.NewEncoder(os.Stdout).Encode(req.Redacted())
	json.NewEncoder(os.Stdout).Encode(req)
}

type SignInParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r SignInParams) Redacted() any {
	r.Password = "********"
	return r
}

// END OMIT
