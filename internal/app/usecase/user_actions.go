package usecase

type IUseractions interface {
	FetchFollowing()
	FetchFollowers(count *int)
	Unfollow()
	Follow()
}

func GetFollows(u IUseractions, count *int) {
	u.FetchFollowers(count)
	u.FetchFollowing()
}

func Follow(u IUseractions, username string) {
	u.Follow()
}
