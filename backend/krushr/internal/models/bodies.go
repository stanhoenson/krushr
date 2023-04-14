package models

type PostRouteBody struct {
	Name     string `json:"name" binding:"required"`
	StatusID uint   `json:"status_id" binding:"required"`
}
