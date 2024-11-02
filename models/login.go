package models

type LoginRes struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
