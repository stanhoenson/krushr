package models

type PostRouteBody struct {
	Title    string `json:"title" binding:"required"`
	StatusID uint   `json:"status_id" binding:"required"`
}
