package models

type User struct {
	ID        int    `json:"id"`
	FULLName  string `json:"full_name"`
	Age       int    `json:"age"`
	AvatarURL string `json: "avatar_url"`
	IsAdmin   bool   `json:"is_admin"`
}
