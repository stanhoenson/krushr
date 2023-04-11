package services

import (
	"errors"
	"time"

	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserFromSignUpBody(signUpBody *models.SignUpBody) (*models.User, error) {
	role, err := repositories.GetEntity[models.Role](env.DefaultRoleID)
	if err != nil {
		return nil, err
	}

	var user models.User
	user.Email = signUpBody.Email
	user.Role = *role

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(signUpBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(passwordBytes)

	return repositories.CreateEntity(&user)
}

func Authenticate(signInBody *models.SignInBody) (string, error) {
	user, err := repositories.GetUserByEmail(signInBody.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInBody.Password))

	if err != nil {
		return "", err
	}

	dayInHours := time.Duration(24)
	weekInDays := time.Duration(7)
	token, err := GenerateJWTWithUser(user, dayInHours*weekInDays)

	return token, nil
}

type CustomClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func GenerateJWTWithUser(user *models.User, hoursValid time.Duration) (string, error) {
	claims := CustomClaims{
		UserID:         user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * hoursValid).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env.JWTSecret))
}

func GetUserFromJWT(jwtString string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(jwtString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(env.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("Invalid claims")
	}

	user, err := repositories.GetUserByIDWithRole(claims.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
