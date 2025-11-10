package domain

type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Hashed_password string `json:"-"`
	Points          int    `json:"points"`
	ReferrerID      *int   `json:"referrer_id"`
}
