package entity

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Created  int64  `json:"created"`
	Modified int64  `json:"modified"`
	Deleted  int64  `json:"deleted"`
}
