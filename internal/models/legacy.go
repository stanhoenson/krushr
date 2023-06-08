package models

type LegacyRoute struct {
	RouteID          uint                    `json:"RouteId"`
	RouteName        string                  `json:"RouteName"`
	Description      string                  `json:"Description"`
	RouteImage       string                  `json:"RouteImage"`
	RouteLength      string                  `json:"RouteLength"`
	RouteType        string                  `json:"RouteType"`
	HasInternalImage bool                    `json:"HasInternalImage"`
	Menu             LegacyMenu              `json:"Menu"`
	RouteOrder       uint                    `json:"RouteOrder"`
	StatusID         uint                    `json:"StatusId"`
	Languages        []string                `json:"Languages"`
	POIList          []LegacyPointOfInterest `json:"POIList"`
}

type LegacyMenu struct {
	MenuID          uint   `json:"MenuId"`
	MenuName        string `json:"MenuName"`
	MenuOrder       uint   `json:"MenuOrder"`
	MenuIcon        string `json:"MenuIcon"`
	MenuDisplayName string `json:"MenuDisplayName"`
}

type LegacyPointOfInterest struct {
	POIID        uint             `json:"POIId"`
	POIName      string           `json:"POIName"`
	Longitude    string           `json:"Longitude"`
	Latitude     string           `json:"Latitude"`
	OrderInRoute uint             `json:"OrderInRoute"`
	CategoryList []LegacyCategory `json:"CategoryList"`
	InfoList     []LegacyInfo     `json:"InfoList"`
}

type LegacyCategory struct {
	CategoryID       uint   `json:"CategoryId"`
	CategoryName     string `json:"CategoryName"`
	CategoryImageURL string `json:"CategoryImageURL"`
	Description      string `json:"Description"`
}

type LegacyInfo struct {
	InfoID         uint              `json:"InfoId"`
	InfoURL        string            `json:"InfoUrl"`
	InfoURLText    string            `json:"InfoUrlText"`
	InfoURLAddress string            `json:"InfoUrlAddress"`
	Omschrijving   string            `json:"Omschrijving"`
	InternalText   string            `json:"InternalText"`
	ContentType    LegacyContentType `json:"ContentType"`
}

type LegacyContentType struct {
	ContentTypeID   uint   `json:"ContentTypeId"`
	ContentTypeName string `json:"ContentTypeName"`
}
