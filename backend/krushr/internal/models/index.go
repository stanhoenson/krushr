package models

type Route struct {
	ID               uint              `json:"id"`
	Name             string            `json:"name"`
	Images           []Image           `json:"images"`
	Details          []Detail          `json:"details"`
	Links            []Link            `json:"links"`
	Categories       []Category        `json:"categories"`
	Status           Status            `json:"status"`
	PointsOfInterest []PointOfInterest `json:"points_of_interest"`
	User             User              `json:"user"`
}

type Image struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type Detail struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

type Link struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

type Category struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Position uint   `json:"position"`
}

type Status struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PointOfInterest struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name"`
	Longitude  float64    `json:"longitude"`
	Latitude   float64    `json:"latitude"`
	Images     []Image    `json:"images"`
	Details    []Detail   `json:"details"`
	Links      []Link     `json:"links"`
	Categories []Category `json:"categories"`
	User       User       `json:"user"`
}

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
