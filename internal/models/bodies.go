package models

// Route
type PostRouteBody struct {
	Name             string                    `json:"name" binding:"required"`
	PointsOfInterest []PostPointOfInterestBody `json:"pointsOfInterest" binding:"required,dive,min=2"`
	ImageIDs         []uint                    `json:"imageIds" binding:"required,min=1"`
	Details          []PostDetailBody          `json:"details" binding:"required,dive,min=1"`
	Links            []PostLinkBody            `json:"links" binding:"dive"`
	Categories       []PostCategoryBody        `json:"categories" binding:"dive"`
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

type PutRouteBody = PostRouteBody

// Image, never used
type PostImageBody struct {
	Path string `json:"path" binding:"required"`
}
type PutImageBody struct {
	PostImageBody
}

// Detail
type PostDetailBody struct {
	Text string `json:"text" binding:"required"`
}
type PutDetailBody struct {
	PostDetailBody
}

// Link
type PostLinkBody struct {
	Text string `json:"text" binding:"required"`
	URL  string `json:"url" binding:"required,url"`
}
type PutLinkBody struct {
	PostLinkBody
}

// Category
type PostCategoryBody struct {
	Name     string `json:"name" binding:"required"`
	Position uint   `json:"position"`
}
type PutCategoryBody struct {
	PostCategoryBody
}

// Status
type PostStatusBody struct {
	Name string `json:"name" binding:"required"`
}
type PutStatusBody struct {
	PostStatusBody
}

// PointOfInterest
type PostPointOfInterestBody struct {
	GetPointOfInterestBody
	ImageIDs   []uint             `json:"imageIds"`
	Details    []PostDetailBody   `json:"details" `
	Links      []PostLinkBody     `json:"links" `
	Categories []PostCategoryBody `json:"categories" `
	Support    bool               `json:"support"`
}

type GetPointOfInterestBody struct {
	Name      string  `json:"name" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required,longitude"`
	Latitude  float64 `json:"latitude" binding:"required,latitude"`
}

func (g *GetPointOfInterestBody) ToPointOfInterest() PointOfInterest {
	return PointOfInterest{
		Name:      g.Name,
		Longitude: g.Longitude,
		Latitude:  g.Latitude,
	}
}

func (p *PostPointOfInterestBody) ToGetPointOfInterestBody() GetPointOfInterestBody {
	return GetPointOfInterestBody{
		Name:      p.Name,
		Longitude: p.Longitude,
		Latitude:  p.Latitude,
	}
}

type PutPointOfInterestBody = PostPointOfInterestBody

type PostUserBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,sha256"`
	RoleID   uint   `json:"roleId" binding:"required"`
}

type PutUserBody struct {
	PostUserBody
}
