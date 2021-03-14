package models

// User : type user struct
type User struct {
	Model
	UserID            string `json:"user_id"`
	Password          string `json:"password"`
	Name              string `json:"name"`
	Phone             string `json:"phone"`
	Address           string `json:"address"`
	BackNumber        int    `json:"back_number"`
	PreferredPosition string `json:"preferred_position"`
}
