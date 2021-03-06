package user

//User ... User model
type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" validate:"required" db:"username"`
	Password string `json:"password" db:"password"`
}
