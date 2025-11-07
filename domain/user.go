package domain

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Points     int    `json:"points"`
	ReferrerID int    `json:"referrer_id"`
}
