package model

type CxsjUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}
