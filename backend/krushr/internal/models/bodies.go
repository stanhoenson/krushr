package models

// Route
type PostRouteBody struct {
	Name             string                    `json:"name" binding:"required"`
	PointsOfInterest []PostPointOfInterestBody `json:"pointsOfInterest" binding:"required,min=2"`
	ImageIDs         []uint                    `json:"imageIds"`
	Details          []PostDetailBody          `json:"details"`
	Links            []PostLinkBody            `json:"links"`
	Categories       []PostCategoryBody        `json:"categories"`
	StatusID         uint                      `json:"statusId" binding:"required"`
}
type PostRouteBodyOld struct {
	Name               string `json:"name" binding:"required"`
	PointOfInterestIDs []uint `json:"pointOfInterestIds"`
	ImageIDs           []uint `json:"imageIds"`
	DetailIDs          []uint `json:"detailIds"`
	LinkIDs            []uint `json:"linkIds"`
	CategoryIDs        []uint `json:"categoryIds"`
	StatusID           uint   `json:"statusId" binding:"required"`
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
	Name       string             `json:"name"`
	Longitude  float64            `json:"longitude"`
	Latitude   float64            `json:"latitude"`
	ImageIDs   []uint             `json:"imageIds"`
	Details    []PostDetailBody   `json:"details"`
	Links      []PostLinkBody     `json:"links"`
	Categories []PostCategoryBody `json:"categories"`
}

type PutPointOfInterestBody struct {
	ID uint `json:"id" binding:"required"`
	PostPointOfInterestBody
}
