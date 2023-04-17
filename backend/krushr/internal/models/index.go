package models

type Route struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Name             string             `gorm:"not null" json:"name"`
	Images           []*Image           `gorm:"many2many:routes_images" json:"images"`
	Details          []*Detail          `gorm:"many2many:routes_details" json:"details"`
	Links            []*Link            `gorm:"many2many:routes_links" json:"links"`
	Categories       []*Category        `gorm:"many2many:routes_categories" json:"categories"`
	Status           Status             `json:"status"`
	StatusID         uint               `gorm:"not null" json:"status_id"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:routes_points_of_interest" json:"points_of_interest"`
	User             User               `json:"user"`
	UserID           uint               `gorm:"not null" json:"user_id"`
}

type Image struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Path             string             `gorm:"not null" json:"path"`
	Routes           []*Route           `gorm:"many2many:routes_images" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_images" json:"points_of_interest"`
}

type Detail struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Text             string             `gorm:"not null" json:"text"`
	Routes           []*Route           `gorm:"many2many:routes_images" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_details" json:"points_of_interest"`
}

type Link struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	URL              string             `gorm:"not null" json:"url"`
	Routes           []*Route           `gorm:"many2many:routes_links" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_links" json:"points_of_interest"`
}

type Category struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Name             string             `gorm:"not null" json:"name"`
	Position         uint               `gorm:"not null"  json:"position"`
	Routes           []*Route           `gorm:"many2many:routes_images" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:points_of_interest_categories" json:"points_of_interest"`
}

type Status struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type PointOfInterest struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	Name       string      `gorm:"not null" json:"name"`
	Longitude  float64     `gorm:"not null" json:"longitude"`
	Latitude   float64     `gorm:"not null" json:"latitude"`
	Images     []*Image    `gorm:"many2many:points_of_interest_images" json:"images"`
	Details    []*Detail   `gorm:"many2many:points_of_interest_details" json:"details"`
	Links      []*Link     `gorm:"many2many:points_of_interest_links" json:"links"`
	Categories []*Category `gorm:"many2many:points_of_interest_categories" json:"categories"`
	User       User        `json:"user"`
	UserID     uint        `gorm:"not null" json:"user_id"`
}

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     Role   `json:"role"`
	RoleID   uint   `gorm:"not null" json:"role_id"`
}

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null, unique" json:"name"`
}
