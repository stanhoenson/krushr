package models

// TODO add validation tags https://github.com/go-playground/validator
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	RoleID   uint   `gorm:"not null" json:"role_id"`
	Role     Role   `json:"role"`
}

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Role string `gorm:"not null" json:"role"`
}

type Route struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Title            string             `gorm:"not null" json:"title"`
	StatusID         uint               `gorm:"not null" json:"status_id"`
	Status           Status             `json:"status"`
	UserID           uint               `gorm:"not null" json:"user_id"`
	User             User               `json:"user"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:routes_points_of_interest" json:"points_of_interest"`
	Entries          []*Entry           `gorm:"many2many:entries_routes" json:"entries"`
	Categories       []*Category        `gorm:"many2many:categories_routes" json:"categories"`
}

type Status struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Status string `gorm:"not null" json:"status"`
}

type PointOfInterest struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	Title      string      `gorm:"not null" json:"title"`
	Longitude  float64     `gorm:"type:decimal(9,6)" json:"longitude"`
	Latitude   float64     `gorm:"type:decimal(9,6)" json:"latitude" binding:"required,latitude"`
	Categories []*Category `gorm:"many2many:categories_points_of_interest" json:"categories"`
	Routes     []*Route    `gorm:"many2many:routes_points_of_interest" json:"routes"`
	Entries    []*Entry    `gorm:"many2many:entries_points_of_interest" json:"entries"`
}

type Type struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Type string `gorm:"not null" json:"type"`
	Icon string `json:"icon"`
}

type Entry struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Content          string             `gorm:"not null" json:"content"`
	Hyperlink        string             `gorm:"not null" json:"hyperlink"`
	TypeID           uint               `gorm:"not null" json:"type_id"`
	Type             Type               `json:"type"`
	Routes           []*Route           `gorm:"many2many:entries_routes" json:"routes"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:entries_points_of_interest" json:"points_of_interest"`
}

type Category struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	Category         string             `gorm:"not null" json:"category"`
	Icon             string             `gorm:"not null" json:"icon"`
	Weight           int                `gorm:"not null" json:"weight"`
	TypeID           uint               `json:"type_id"`
	Type             Type               `json:"type"`
	PointsOfInterest []*PointOfInterest `gorm:"many2many:categories_points_of_interest" json:"points_of_interest"`
	Routes           []*Route           `gorm:"many2many:categories_routes" json:"routes"`
}
