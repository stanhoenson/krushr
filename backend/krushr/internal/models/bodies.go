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
	PostRouteBody
}

// Image
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
	URL string `json:"url" binding:"required"`
}
type PutLinkBody struct {
	PostLinkBody
}

// Category
type PostCategoryBody struct {
	Name     string `json:"name" binding:"required"`
	Position uint   `json:"position" binding:"required"`
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
	Details    []PostDetailBody   `json:"details"`
	Links      []PostLinkBody     `json:"links"`
	Categories []PostCategoryBody `json:"categories"`
}

type GetPointOfInterestBody struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
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

type PutPointOfInterestBody struct {
	PostPointOfInterestBody
}
