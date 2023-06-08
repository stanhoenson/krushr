package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(postUserBody *models.PostUserBody, tx *gorm.DB) (*models.User, error) {
	role, err := repositories.GetEntityByID[models.Role](postUserBody.RoleID, tx)
	if err != nil {
		return nil, err
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(postUserBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:    postUserBody.Email,
		Password: string(passwordBytes),
		Role:     *role,
	}

	createdUser, err := repositories.CreateEntity(&user, tx)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func UpdateUser(ID uint, putUserBody *models.PutUserBody, tx *gorm.DB) (*models.User, error) {
	user, err := repositories.GetEntityByID[models.User](ID, tx)
	if err != nil {
		return nil, err
	}

	role, err := repositories.GetEntityByID[models.Role](putUserBody.RoleID, tx)
	if err != nil {
		return nil, err
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(putUserBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Email = putUserBody.Email
	user.Password = string(passwordBytes)
	user.Role = *role

	updatedUser, err := repositories.UpdateEntity(user, tx)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
