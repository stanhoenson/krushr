package services

import (
	"errors"
	"time"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserFromSignUpBody(signUpBody *models.SignUpBody) (*models.User, error) {
	role, err := repositories.GetEntityByID[models.Role](env.DefaultRoleID, database.Db)
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

	return repositories.CreateEntity(&user, database.Db)
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

	expirationTime := time.Now().Add(constants.TokenValidityPeriod)
	token, err := GenerateJWTWithUser(user, expirationTime)
	if err != nil {
		return "", err
	}

	return token, nil
}

type CustomClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func GenerateJWTWithUser(user *models.User, expirationTime time.Time) (string, error) {
	claims := CustomClaims{
		UserID:         user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
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
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	user, err := repositories.GetUserByIDWithRole(claims.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
