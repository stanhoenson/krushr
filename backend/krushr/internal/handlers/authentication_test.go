package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/middleware"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationRoutes(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.Authorization())
	handlers.RegisterAuthenticationRoutes(r)
	database.InitializeDatabase("test.db", "test/")
	populateDatabaseWithDummyData()

	t.Run("authentication", func(t *testing.T) {
		t.Run("testSignIn", func(t *testing.T) {
			testSignIn(t, r)
		})
	})

	os.Remove("test/test.db")
	os.Remove("test")
}

func testSignIn(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{Email: user.Email, Password: utils.Sha256(env.AdminPassword)}

	signInBodyJson, _ := json.Marshal(signInBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/authentication/sign-in", bytes.NewBuffer(signInBodyJson))
	r.ServeHTTP(w, req)

	responseCookies := w.Result().Cookies()

	jwtCookieFound := false
	for _, cookie := range responseCookies {
		if cookie.Name == "jwt" {
			jwtCookieFound = true
			token, err := jwt.ParseWithClaims(cookie.Value, &services.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(env.JWTSecret), nil
			})
			if err != nil || !token.Valid {
				t.Errorf("faulty jwt, token invalid")
			}
			claims, ok := token.Claims.(*services.CustomClaims)
			if !ok {
				t.Errorf("faulty jwt, no custom claims")
			}
			if user.ID != claims.UserID {
				t.Errorf("faulty jwt, incorrect userID")
			}

		}
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, true, jwtCookieFound)
}
