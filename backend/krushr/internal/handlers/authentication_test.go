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
	env.InitializeEnvironment("../../.env")
	handlers.RegisterAuthenticationRoutes(r)
	database.InitializeDatabase("test.db", "test/")

	t.Run("authentication", func(t *testing.T) {
		t.Run("testSignIn", func(t *testing.T) {
			testSignIn(t, r)
		})
		t.Run("testSignInWrongPassword", func(t *testing.T) {
			testSignInWrongPassword(t, r)
		})
		t.Run("testSignInFaultyBody", func(t *testing.T) {
			testSignInFaultyBody(t, r)
		})
		t.Run("testSignUpUnauthorized", func(t *testing.T) {
			testSignUpUnauthorized(t, r)
		})
		t.Run("testSignUp", func(t *testing.T) {
			testSignUp(t, r)
		})
		t.Run("testSignOut", func(t *testing.T) {
			testSignOut(t, r)
		})
		t.Run("testSignOutUnauthenticated", func(t *testing.T) {
			testSignOutUnauthenticated(t, r)
		})
		t.Run("testSignUpFaultyBody", func(t *testing.T) {
			testSignUpFaultyBody(t, r)
		})
		t.Run("testSignUpDuplicateEmail", func(t *testing.T) {
			testSignUpDuplicateEmail(t, r)
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

func testSignInWrongPassword(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{Email: user.Email, Password: utils.Sha256("wrong!!!")}

	signInBodyJson, _ := json.Marshal(signInBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/authentication/sign-in", bytes.NewBuffer(signInBodyJson))
	r.ServeHTTP(w, req)

	responseCookies := w.Result().Cookies()

	jwtCookieFound := false
	for _, cookie := range responseCookies {
		if cookie.Name == "jwt" {
			jwtCookieFound = true
		}
	}

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, false, jwtCookieFound)
}

func testSignInFaultyBody(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{Email: user.Email, Password: "this is not sha"}

	signInBodyJson, _ := json.Marshal(signInBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/authentication/sign-in", bytes.NewBuffer(signInBodyJson))
	r.ServeHTTP(w, req)

	var errorMsg map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &errorMsg)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Key: 'SignInBody.Password' Error:Field validation for 'Password' failed on the 'sha256' tag", errorMsg["error"])
}

func testSignUpUnauthorized(t *testing.T, r *gin.Engine) {
	email := "creator@creator.com"
	signUpBody := models.SignUpBody{Email: email, Password: utils.Sha256(env.AdminPassword)}

	signUpBodyJson, _ := json.Marshal(signUpBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/authentication/sign-up", bytes.NewBuffer(signUpBodyJson))
	r.ServeHTTP(w, req)

	var createdUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &createdUser)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func testSignUp(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}

	email := "creator@creator.com"
	signUpBody := models.SignUpBody{Email: email, Password: utils.Sha256(env.AdminPassword)}

	signUpBodyJson, _ := json.Marshal(signUpBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/authentication/sign-up", bytes.NewBuffer(signUpBodyJson))
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var createdUser models.User
	err = json.Unmarshal(w.Body.Bytes(), &createdUser)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, email, createdUser.Email)
}

func testSignUpFaultyBody(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}

	email := "admin"
	signUpBody := models.SignUpBody{Email: email, Password: utils.Sha256(env.AdminPassword)}

	signUpBodyJson, _ := json.Marshal(signUpBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/authentication/sign-up", bytes.NewBuffer(signUpBodyJson))
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var errorMsg map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &errorMsg)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Key: 'SignUpBody.Email' Error:Field validation for 'Email' failed on the 'email' tag", errorMsg["error"])
}

func testSignUpDuplicateEmail(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}

	email := "admin@admin.com"
	signUpBody := models.SignUpBody{Email: email, Password: utils.Sha256(env.AdminPassword)}

	signUpBodyJson, _ := json.Marshal(signUpBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/authentication/sign-up", bytes.NewBuffer(signUpBodyJson))
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var errorMsg map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &errorMsg)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Error signing up", errorMsg["error"])
}

func testSignOut(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/authentication/sign-out", nil)
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	responseCookies := w.Result().Cookies()
	jwtCookieFound := false
	for _, cookie := range responseCookies {
		if cookie.Name == "jwt" && cookie.Value != "" {
			jwtCookieFound = true
		}
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, false, jwtCookieFound)
}

func testSignOutUnauthenticated(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/authentication/sign-out", nil)
	r.ServeHTTP(w, req)

	responseCookies := w.Result().Cookies()
	jwtCookieFound := false
	for _, cookie := range responseCookies {
		if cookie.Name == "jwt" && cookie.Value != "" {
			jwtCookieFound = true
		}
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, false, jwtCookieFound)
}
