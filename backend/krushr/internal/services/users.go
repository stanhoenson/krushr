package services

import (
	"os"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() []models.User {
	users := repositories.GetUsers()
	return users
}

func CreateUser(user *models.User) (*models.User, error) {
	roleIdString := os.Getenv("DEFAULT_ROLE_ID")

	u64, err := strconv.ParseUint(roleIdString, 10, 64)
	if err != nil {
		return nil, err
	}
	roleID := uint(u64)
	role, err := repositories.GetEntity[models.Role](roleID)
	if err != nil {
		return nil, err
	}
	user.Role = *role

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(passwordBytes)

	return repositories.CreateEntity(user)
}
