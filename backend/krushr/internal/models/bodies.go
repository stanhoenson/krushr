package models

type PostRouteBody struct {
	Name     string `json:"name" binding:"required"`
	StatusID uint   `json:"status_id" binding:"required"`
}

type PutRouteBody struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" binding:"required"`
	StatusID uint   `json:"status_id" binding:"required"`
}

type PostImageBody struct {
	Path string `json:"path" binding:"required"`
}

type PostDetailBody struct {
	Text string `json:"text" binding:"required"`
}

type PostLinkBody struct {
	URL string `json:"url" binding:"required"`
}

type PostCategoryBody struct {
	Name     string `json:"name" binding:"required"`
	Position uint   `json:"position" binding:"required"`
}
type PostStatusBody struct {
	Name string `json:"name" binding:"required"`
}
