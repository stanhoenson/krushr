package models

type Route struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	StatusID string `json:"status_id"`
	UserID   string `json:"user_id"`
}
