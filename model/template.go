package model

type Template struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}
