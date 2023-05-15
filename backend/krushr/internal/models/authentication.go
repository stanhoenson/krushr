package models

type SignUpBody struct {
	Email    string `gorm:"not null" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" binding:"required,sha256"`
}

type SignInBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,sha256"`
}
