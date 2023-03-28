package models

type Route struct {
	ID       string `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	StatusID string `json:"status_id"`
	UserID   string `json:"user_id"`
}
