package model

// Main user struct
type MyUser struct {
	Followers  []User
	Following  []User
	Allowed    []User
	Denied     []User
	Login      string `json:"login"`
	TargetUser string
	token      string
}

// a single user struct
type User struct {
	Login string `json:"login"`
}

func (u *MyUser) FetchFollowing() {
	//TODO: Implement this function
}

func (u *MyUser) FetchFollowers(count *int) {
	//TODO: Implement this function
}

func (u *MyUser) Unfollow() {
	//TODO: Implement this function
}

func (u *MyUser) Follow() {
	//TODO: Implement this function
}
