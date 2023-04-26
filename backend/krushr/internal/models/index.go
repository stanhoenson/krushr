package models

import (
	"fmt"
	"strconv"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/env"
)

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

func (r Route) ToLegacyRoute() (*LegacyRoute, error) {
	distance := strconv.FormatFloat(r.Distance, 'f', 2, 64)
	if len(r.Images) == 0 {
		return nil, fmt.Errorf("No image for route")
	}
	firstImage := r.Images[0]
	if len(r.Categories) == 0 {
		return nil, fmt.Errorf("No category for route")
	}
	firstCategory := r.Categories[0]

	var poiList []LegacyPointOfInterest

	for _, v := range r.PointsOfInterest {
		poiList = append(poiList, v.ToLegacyPointOfInterest())
	}

	var description string
	for index, v := range r.Details {
		if index == 0 {

			description += v.Text
		} else {

			description += "\n\n" + v.Text
		}

	}

	return &LegacyRoute{
		RouteID:          r.ID,
		Description:      description,
		RouteName:        r.Name,
		RouteImage:       fmt.Sprint(env.ApiUrl, "imagedata/", firstImage.ID),
		RouteLength:      distance,
		RouteType:        firstCategory.Name,
		HasInternalImage: true,
		Menu:             firstCategory.ToLegacyMenu(), RouteOrder: 0,
		Languages: constants.DefaultLegacyRouteLanguages,
		POIList:   poiList,
	}, nil
}

type Image struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Path             string             `gorm:"not null;unique" json:"path"`
	Routes           []*Route           `gorm:"many2many:routes_images;constraint:OnDelete:CASCADE" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_images" json:"pointsOfInterest"`
}

type Detail struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Text             string             `gorm:"not null;unique" json:"text"`
	Routes           []*Route           `gorm:"many2many:routes_details;constraint:OnDelete:CASCADE" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_details;constraint:OnDelete:CASCADE" json:"pointsOfInterest"`
}

type Link struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	URL              string             `gorm:"not null;unique" json:"url"`
	Routes           []*Route           `gorm:"many2many:routes_links;constraint:OnDelete:CASCADE" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_links;constraint:OnDelete:CASCADE" json:"pointsOfInterest"`
}

type Category struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Name             string             `gorm:"not null;unique" json:"name"`
	Position         uint               `gorm:"not null"  json:"position"`
	Routes           []*Route           `gorm:"many2many:routes_categories;constraint:OnDelete:CASCADE" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_categories;constraint:OnDelete:CASCADE" json:"pointsOfInterest"`
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
}

func (p PointOfInterest) ToLegacyPointOfInterest() LegacyPointOfInterest {
	return LegacyPointOfInterest{}
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

type RoutesPointsOfInterest struct {
	RouteID           uint `gorm:"primaryKey"`
	PointOfInterestID uint `gorm:"primaryKey"`
	Position          uint `gorm:"not null"`
}

func (RoutesPointsOfInterest) TableName() string {
	return "routes_points_of_interest"
}
