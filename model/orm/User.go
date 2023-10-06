package orm

type User struct {
	Uid      int
	Username string
	Password string
	Salt     string
	Rating   int
}
