package model

// Main user struct
type MyUser struct {
	Followers []User
	Folliwng  []User
	Allowed   []User
	Denied    []User
	Login     string `json:"login"`
	token     string
}

// a single user struct
type User struct {
	Login string `json:"login"`
}
