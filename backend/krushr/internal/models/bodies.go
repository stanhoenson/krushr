package models

// Route
type PostRouteBody struct {
	Name     string `json:"name" binding:"required"`
	StatusID uint   `json:"status_id" binding:"required"`
}

type PutRouteBody struct {
	ID uint `json:"id" binding:"required"`
	PostRouteBody
}

// Image
type PostImageBody struct {
	Path string `json:"path" binding:"required"`
}
type PutImageBody struct {
	ID uint `json:"id" binding:"required"`
	PostImageBody
}

// Detail
type PostDetailBody struct {
	Text string `json:"text" binding:"required"`
}
type PutDetailBody struct {
	ID uint `json:"id" binding:"required"`
	PostDetailBody
}

// Link
type PostLinkBody struct {
	URL string `json:"url" binding:"required"`
}
type PutLinkBody struct {
	ID uint `json:"id" binding:"required"`
	PostLinkBody
}

// Category
type PostCategoryBody struct {
	Name     string `json:"name" binding:"required"`
	Position uint   `json:"position" binding:"required"`
}
type PutCategoryBody struct {
	ID uint `json:"id" binding:"required"`
	PostCategoryBody
}

// Status
type PostStatusBody struct {
	Name string `json:"name" binding:"required"`
}
type PutStatusBody struct {
	ID uint `json:"id" binding:"required"`
	PostStatusBody
}

// PointOfInterest
type PostPointOfInterestBody struct {
	Name        string      `json:"name"`
	Longitude   float64     `json:"longitude"`
	Latitude    float64     `json:"latitude"`
	ImagesIDs   []*Image    `json:"images"`
	DetailIDs   []*Detail   `json:"details"`
	LinkIDs     []*Link     `json:"links"`
	CategoryIDs []*Category `json:"categories"`
}

type PutPointOfInterestBody struct {
	ID uint `json:"id" binding:"required"`
	PostPointOfInterestBody
}
