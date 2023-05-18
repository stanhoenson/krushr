package models

import (
	"fmt"
	"strconv"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/env"
	"gorm.io/gorm"
)

func checkAndDeleteRelatedEntity[T any, B any](tx *gorm.DB, whereStatement string, entityTwoID uint) error {
	var entityOne T
	var entityTwo B
	// Check if the related entity (Entity2) has any other relationships
	var count int64
	if err := tx.Model(&entityOne).Where(whereStatement, entityTwoID).Count(&count).Error; err != nil {
		return err
	}
	fmt.Println(count)

	// If there are no other relationships, delete the related entity (Entity2)
	if count == 0 {
		if err := tx.Delete(&entityTwo, entityTwoID).Error; err != nil {
			return err
		}
	}

	return nil
}

type Route struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Name             string             `gorm:"not null;unique" json:"name"`
	Images           []*Image           `gorm:"many2many:routes_images;constraint:OnDelete:CASCADE" json:"images"`
	Details          []*Detail          `gorm:"many2many:routes_details;constraint:OnDelete:CASCADE" json:"details"`
	Links            []*Link            `gorm:"many2many:routes_links;constraint:OnDelete:CASCADE" json:"links"`
	Categories       []*Category        `gorm:"many2many:routes_categories;constraint:OnDelete:CASCADE" json:"categories"`
	Status           Status             `json:"status"`
	StatusID         uint               `gorm:"not null" json:"statusId"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:routes_points_of_interest;constraint:OnDelete:CASCADE" json:"pointsOfInterest"`
	Distance         float64            `gorm:"not null" json:"distance"`
	User             User               `json:"user"`
	UserID           uint               `gorm:"not null" json:"userId"`
}

type RouteImage struct {
	RouteID uint `gorm:"primaryKey"`
	ImageID uint `gorm:"primaryKey"`
}

func (RouteImage) TableName() string {
	return "routes_images"
}

func (e *RouteImage) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("in the function man")
	fmt.Println(e)
	fmt.Println(tx.Statement)
	if err := checkAndDeleteRelatedEntity[RouteImage, Image](tx, "image_id = ?", e.ImageID); err != nil {
		return err
	}
	return nil
}

type RouteDetail struct {
	RouteID  uint `gorm:"primaryKey"`
	DetailID uint `gorm:"primaryKey"`
}

type RouteLink struct {
	RouteID uint `gorm:"primaryKey"`
	LinkID  uint `gorm:"primaryKey"`
}

type RouteCategory struct {
	RouteID    uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
}

type RoutesPointsOfInterest struct {
	RouteID           uint `gorm:"primaryKey"`
	PointOfInterestID uint `gorm:"primaryKey"`
	Position          uint `gorm:"not null"`
}

func (RoutesPointsOfInterest) TableName() string {
	return "routes_points_of_interest"
}

func (r Route) ToLegacyRoute(withPOIs bool) (*LegacyRoute, error) {
	distance := strconv.FormatFloat(r.Distance, 'f', 2, 64)
	if len(r.Images) == 0 {
		return nil, fmt.Errorf("no image for route")
	}
	firstImage := r.Images[0]
	if len(r.Categories) == 0 {
		return nil, fmt.Errorf("no category for route")
	}
	firstCategory := r.Categories[0]

	var description string
	for index, v := range r.Details {
		if index == 0 {
			description += v.Text
		} else {
			description += "\n\n" + v.Text
		}
	}

	legacyRoute := &LegacyRoute{
		RouteID:          r.ID,
		Description:      description,
		RouteName:        r.Name,
		RouteImage:       fmt.Sprint(env.ApiUrl, "imagedata/", firstImage.ID),
		RouteLength:      distance,
		RouteType:        firstCategory.Name,
		HasInternalImage: true,
		Menu:             firstCategory.ToLegacyMenu(),
		RouteOrder:       r.ID,
		Languages:        constants.DefaultLegacyRouteLanguages,
	}
	if withPOIs {

		var poiList []LegacyPointOfInterest

		for index, v := range r.PointsOfInterest {
			poiList = append(poiList, v.ToLegacyPointOfInterest(uint(index)))
		}
		legacyRoute.POIList = poiList
	}
	return legacyRoute, nil
}

type Image struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Path             string             `gorm:"not null;unique" json:"path"`
	Routes           []*Route           `gorm:"many2many:routes_images;constraint:OnDelete:CASCADE" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_images" json:"pointsOfInterest"`
}

func (i Image) ToLegacyInfo() LegacyInfo {
	return LegacyInfo{
		InfoID: i.ID,
		ContentType: LegacyContentType{
			ContentTypeID:   constants.LegacyImageContentTypeId,
			ContentTypeName: constants.LegacyImageContentTypeName,
		},
	}
}

type Detail struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Text             string             `gorm:"not null;unique" json:"text"`
	Routes           []*Route           `gorm:"many2many:routes_details;constraint:OnDelete:CASCADE" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_details;constraint:OnDelete:CASCADE" json:"pointsOfInterest"`
}

func (d Detail) ToLegacyInfo() LegacyInfo {
	return LegacyInfo{
		InfoID:       d.ID,
		Omschrijving: d.Text,
		InternalText: d.Text,
		ContentType: LegacyContentType{
			ContentTypeID:   constants.LegacyTekstContentTypeId,
			ContentTypeName: constants.LegacyTekstContentTypeName,
		},
	}
}

type Link struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Text             string             `gorm:"not null" json:"text"`
	URL              string             `gorm:"not null;unique" json:"url"`
	Routes           []*Route           `gorm:"many2many:routes_links;constraint:OnDelete:CASCADE" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_links;constraint:OnDelete:CASCADE" json:"pointsOfInterest"`
}

func (l Link) ToLegacyInfo() LegacyInfo {
	return LegacyInfo{
		InfoID:  l.ID,
		InfoURL: l.URL,
		ContentType: LegacyContentType{
			ContentTypeID:   constants.LegacyWebsiteContentTypeId,
			ContentTypeName: constants.LegacyWebsiteContentTypeName,
		},
	}
}

type Category struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Name             string             `gorm:"not null;unique" json:"name"`
	Position         uint               `gorm:"not null"  json:"position"`
	Routes           []*Route           `gorm:"many2many:routes_categories;constraint:OnDelete:CASCADE" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_categories;constraint:OnDelete:CASCADE" json:"pointsOfInterest"`
}

func (c Category) ToLegacyCategory() LegacyCategory {
	return LegacyCategory{
		CategoryID:       c.ID,
		CategoryName:     c.Name,
		CategoryImageURL: "Geen URL",
		Description:      c.Name,
	}
}

func (c Category) ToLegacyMenu() LegacyMenu {
	return LegacyMenu{
		MenuID:          c.ID,
		MenuName:        c.Name,
		MenuOrder:       c.Position,
		MenuIcon:        "",
		MenuDisplayName: c.Name,
	}
}

type Status struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null;unique" json:"name"`
}

type PointOfInterest struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	Name       string      `gorm:"not null;uniqueIndex:poiIndex" json:"name"`
	Longitude  float64     `gorm:"not null;uniqueIndex:poiIndex" json:"longitude"`
	Latitude   float64     `gorm:"not null;uniqueIndex:poiIndex" json:"latitude"`
	Images     []*Image    `gorm:"many2many:points_of_interest_images;constraint:OnDelete:CASCADE" json:"images"`
	Details    []*Detail   `gorm:"many2many:points_of_interest_details;constraint:OnDelete:CASCADE" json:"details"`
	Links      []*Link     `gorm:"many2many:points_of_interest_links;constraint:OnDelete:CASCADE" json:"links"`
	Categories []*Category `gorm:"many2many:points_of_interest_categories;constraint:OnDelete:CASCADE" json:"categories"`
	Routes     []*Route    `gorm:"many2many:routes_points_of_interest;constraint:OnDelete:CASCADE" json:"routes"`
	User       User        `json:"user"`
	UserID     uint        `gorm:"not null" json:"userId"`
	Support    bool        `gorm:"not null;default:false" json:"support"`
}

type PointOfInterestImage struct {
	PointOfInterestID uint `gorm:"primaryKey"`
	ImageID           uint `gorm:"primaryKey"`
}

type PointOfInterestDetail struct {
	PointOfInterestID uint `gorm:"primaryKey"`
	DetailID          uint `gorm:"primaryKey"`
}

type PointOfInterestLink struct {
	PointOfInterestID uint `gorm:"primaryKey"`
	LinkID            uint `gorm:"primaryKey"`
}

type PointOfInterestCategory struct {
	PointOfInterestID uint `gorm:"primaryKey"`
	CategoryID        uint `gorm:"primaryKey"`
}

func (p PointOfInterest) ToLegacyPointOfInterest(orderInRoute uint) LegacyPointOfInterest {
	longitude := strconv.FormatFloat(p.Longitude, 'f', 2, 64)
	latitude := strconv.FormatFloat(p.Latitude, 'f', 2, 64)
	var categoryList []LegacyCategory
	for _, v := range p.Categories {
		categoryList = append(categoryList, v.ToLegacyCategory())
	}

	var infoList []LegacyInfo

	for _, v := range p.Details {
		infoList = append(infoList, v.ToLegacyInfo())
	}
	for _, v := range p.Links {
		infoList = append(infoList, v.ToLegacyInfo())
	}
	for _, v := range p.Images {
		infoList = append(infoList, v.ToLegacyInfo())
	}

	return LegacyPointOfInterest{
		POIID:        p.ID,
		POIName:      p.Name,
		Longitude:    longitude,
		Latitude:     latitude,
		CategoryList: categoryList,
		InfoList:     infoList,
		OrderInRoute: orderInRoute,
	}
}

func (PointOfInterest) TableName() string {
	return "points_of_interest"
}

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role     Role   `json:"role"`
	RoleID   uint   `gorm:"not null" json:"roleId"`
}

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null;unique" json:"name"`
}
