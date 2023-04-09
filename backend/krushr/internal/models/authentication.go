package models

type SignUpBody struct {
	Email    string `gorm:"not null" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" binding:"required"`
}

type SignInBody struct {
	Email    string `gorm:"not null" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" binding:"required"`
}
